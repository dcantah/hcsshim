package jobs

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"unsafe"

	"github.com/Microsoft/hcsshim/internal/log"
	"github.com/Microsoft/hcsshim/internal/winapi"
	"golang.org/x/sys/windows"
)

// This file provides higher level constructs for the win32 job object API.
// Most of the core creation and management functions are already present in "golang.org/x/sys/windows"
// (CreateJobObject, AssignProcessToJobObject, etc.) as well as most of the limit information
// structs and associated limit flags. Whatever is not present from the job object API
// in golang.org/x/sys/windows is located in /internal/winapi.
//
// https://docs.microsoft.com/en-us/windows/win32/procthread/job-objects

// JobObject is a high level wrapper around a Windows job object. Holds a handle to
// the job, a handle to an iocp to be used to receive notifications about the lifecycle
// of the job and a mutex for synchronized handle access.
type JobObject struct {
	jobHandle  windows.Handle
	iocpChan   chan IOCPNotif
	handleLock sync.RWMutex
}

type jobLimits struct {
	affinity       uintptr
	cpuRate        uint32
	cpuWeight      uint32
	jobMemoryLimit uintptr
	maxIops        int64
	maxBandwidth   int64
}

var errAlreadyClosed = errors.New("the handle has already been closed")

// SetResourceLimits sets resource limits on the job object (cpu, memory, storage).
func (job *JobObject) SetResourceLimits(ctx context.Context, limits *jobLimits) error {
	if job.jobHandle == 0 {
		return errAlreadyClosed
	}
	job.handleLock.RLock()
	defer job.handleLock.RUnlock()

	// Go through and check what limits were specified and construct the appropriate
	// structs.
	if limits.affinity != 0 || limits.jobMemoryLimit != 0 {
		var (
			basicLimitFlags uint32
			eliInfo         windows.JOBOBJECT_EXTENDED_LIMIT_INFORMATION
		)
		if limits.affinity != 0 {
			basicLimitFlags |= windows.JOB_OBJECT_LIMIT_AFFINITY
			eliInfo.BasicLimitInformation.Affinity = limits.affinity
		}
		if limits.jobMemoryLimit != 0 {
			basicLimitFlags |= windows.JOB_OBJECT_LIMIT_JOB_MEMORY
			eliInfo.JobMemoryLimit = limits.jobMemoryLimit
		}
		eliInfo.BasicLimitInformation.LimitFlags = basicLimitFlags
		_, err := windows.SetInformationJobObject(job.jobHandle, windows.JobObjectExtendedLimitInformation, uintptr(unsafe.Pointer(&eliInfo)), uint32(unsafe.Sizeof(eliInfo)))
		if err != nil {
			return fmt.Errorf("failed to set extended limit info on job object: %s", err)
		}
	}

	if limits.cpuRate != 0 {
		cpuInfo := winapi.JOBOBJECT_CPU_RATE_CONTROL_INFORMATION{
			ControlFlags: winapi.JOB_OBJECT_CPU_RATE_CONTROL_ENABLE | winapi.JOB_OBJECT_CPU_RATE_CONTROL_HARD_CAP,
			Rate:         limits.cpuRate,
		}
		_, err := windows.SetInformationJobObject(job.jobHandle, windows.JobObjectCpuRateControlInformation, uintptr(unsafe.Pointer(&cpuInfo)), uint32(unsafe.Sizeof(cpuInfo)))
		if err != nil {
			return fmt.Errorf("failed to set cpu limit info on job object: %s", err)
		}
	}

	if limits.maxBandwidth != 0 || limits.maxIops != 0 {
		ioInfo := winapi.JOBOBJECT_IO_RATE_CONTROL_INFORMATION{
			ControlFlags: winapi.JOB_OBJECT_IO_RATE_CONTROL_ENABLE,
		}
		if limits.maxBandwidth != 0 {
			ioInfo.MaxBandwidth = limits.maxBandwidth
		}
		if limits.maxIops != 0 {
			ioInfo.MaxIops = limits.maxIops
		}
		_, err := winapi.SetIoRateControlInformationJobObject(job.jobHandle, &ioInfo)
		if err != nil {
			return fmt.Errorf("failed to set IO limit info on job object: %s", err)
		}
	}
	return nil
}

// CreateJobObject creates a job object, attaches an IO completion port to use
// for notifications and then returns an object with the corresponding handles.
func CreateJobObject(name string) (*JobObject, error) {
	jobHandle, err := windows.CreateJobObject(nil, windows.StringToUTF16Ptr(name))
	if err != nil {
		return nil, err
	}

	// Job object creation succeeded, add the handle value to the map so we can
	// receive IOCP notifications.
	iocpChan := make(chan IOCPNotif)
	jobMap[uint32(jobHandle)] = iocpChan
	if ioCompletionPort == 0 {
		fmt.Println("Doing this once")
		ioInitOnce.Do(initIo)
	}

	if _, err = attachIOCP(jobHandle, ioCompletionPort); err != nil {
		windows.Close(jobHandle)
		return nil, err
	}
	return &JobObject{
		jobHandle,
		iocpChan,
		sync.RWMutex{},
	}, nil
}

// IOCPChan returns the channel that the job object will receive IO completion port
// notifications on.
func (job *JobObject) IOCPChan() chan IOCPNotif {
	return job.iocpChan
}

// Close closes the job object and iocp handles. If this is the only open handle
// the job object will be terminated.
func (job *JobObject) Close() error {
	if job.jobHandle == 0 {
		return nil
	}
	job.handleLock.Lock()
	defer job.handleLock.Unlock()

	if job.jobHandle != 0 {
		if err := windows.Close(job.jobHandle); err != nil {
			return err
		}
		job.jobHandle = 0
	}
	return nil
}

// Assign assigns a process to the job object.
func (job *JobObject) Assign(pid uint32) error {
	if job.jobHandle == 0 {
		return errAlreadyClosed
	}
	job.handleLock.RLock()
	defer job.handleLock.RUnlock()

	if pid == 0 {
		return errors.New("process has not started")
	}
	hProc, err := windows.OpenProcess(winapi.PROCESS_ALL_ACCESS, true, pid)
	if err != nil {
		return err
	}
	defer windows.Close(hProc)
	return windows.AssignProcessToJobObject(job.jobHandle, hProc)
}

// Terminate terminates the job, essentially calls TerminateProcess on every process in the
// job.
func (job *JobObject) Terminate() error {
	if job.jobHandle == 0 {
		return errAlreadyClosed
	}
	job.handleLock.RLock()
	defer job.handleLock.RUnlock()
	return windows.TerminateJobObject(job.jobHandle, 1)
}

// Shutdown gracefully shuts down all the processes in the job object.
func (job *JobObject) Shutdown(ctx context.Context) error {
	if job.jobHandle == 0 {
		return errAlreadyClosed
	}
	job.handleLock.RLock()
	defer job.handleLock.RUnlock()

	var (
		terminate bool
		signalErr bool
	)
	pids, err := job.Pids()
	if err != nil {
		return fmt.Errorf("failed to get pids for job object: %s", err)
	}

	for _, pid := range pids {
		if err := windows.GenerateConsoleCtrlEvent(windows.CTRL_BREAK_EVENT, pid); err != nil {
			log.G(ctx).WithField("pid", pid).Errorf("failed to send ctrl-break to process in job: %s", err)
			signalErr = true
		}
	}

	// Get pids in job again to check if they were successfully killed by the signals.
	// Certain programs (ping.exe im looking at you) handle ctrl signals so we need
	// a fallback to actually kill them.
	newPids, err := job.Pids()
	if err != nil {
		return fmt.Errorf("failed to get pids for job object: %s", err)
	}
	terminate = len(newPids) != 0
	// If any of the processes couldnt be killed gracefully just terminate the job.
	// Equivalent to calling TerminateProcess on every proc in the job.
	if terminate || signalErr {
		return job.Terminate()
	}
	return nil
}

// Pids returns all of the process IDs in the job object.
func (job *JobObject) Pids() ([]uint32, error) {
	if job.jobHandle == 0 {
		return nil, errAlreadyClosed
	}
	job.handleLock.RLock()
	defer job.handleLock.RUnlock()

	info := winapi.JOBOBJECT_BASIC_PROCESS_ID_LIST{}
	err := winapi.QueryInformationJobObject(
		job.jobHandle,
		winapi.JobObjectBasicProcessIdList,
		uintptr(unsafe.Pointer(&info)),
		uint32(unsafe.Sizeof(info)),
		nil,
	)

	// This is either the case where there is only one process or no processes in
	// the job. Any other case will result in ERROR_MORE_DATA. Check if info.NumberOfProcessIdsInList
	// is 1 and just return this, otherwise return an empty slice.
	if err == nil {
		if info.NumberOfProcessIdsInList == 1 {
			return []uint32{uint32(info.ProcessIdList[0])}, nil
		}
		// Return empty slice instead of nil to play well with the caller of this.
		// Do not return an error if no processes are running inside the job
		return []uint32{}, nil
	}

	if err != winapi.ERROR_MORE_DATA {
		return nil, fmt.Errorf("failed initial query for PIDs in job object: %s", err)
	}

	jobBasicProcessIDListSize := unsafe.Sizeof(info) + (unsafe.Sizeof(info.ProcessIdList[0]) * uintptr(info.NumberOfAssignedProcesses-1))
	buf := make([]byte, jobBasicProcessIDListSize)
	if err = winapi.QueryInformationJobObject(
		job.jobHandle,
		winapi.JobObjectBasicProcessIdList,
		uintptr(unsafe.Pointer(&buf[0])),
		uint32(len(buf)),
		nil,
	); err != nil {
		return nil, fmt.Errorf("failed to query for PIDs in job object: %s", err)
	}

	bufInfo := (*winapi.JOBOBJECT_BASIC_PROCESS_ID_LIST)(unsafe.Pointer(&buf[0]))
	bufPids := bufInfo.AllPids()
	pids := make([]uint32, bufInfo.NumberOfProcessIdsInList)
	for i, bufPid := range bufPids {
		pids[i] = uint32(bufPid)
	}
	return pids, nil
}

// QueryMemoryStats gets the memory stats for the job object.
func (job *JobObject) QueryMemoryStats() (*winapi.JOBOBJECT_MEMORY_USAGE_INFORMATION, error) {
	if job.jobHandle == 0 {
		return nil, errAlreadyClosed
	}
	job.handleLock.RLock()
	defer job.handleLock.RUnlock()

	info := winapi.JOBOBJECT_MEMORY_USAGE_INFORMATION{}
	if err := winapi.QueryInformationJobObject(
		job.jobHandle,
		winapi.JobObjectMemoryUsageInformation,
		uintptr(unsafe.Pointer(&info)),
		uint32(unsafe.Sizeof(info)),
		nil,
	); err != nil {
		return nil, fmt.Errorf("failed to query for job object memory stats: %s", err)
	}
	return &info, nil
}

// QueryProcessorStats gets the processor stats for the job object.
func (job *JobObject) QueryProcessorStats() (*winapi.JOBOBJECT_BASIC_ACCOUNTING_INFORMATION, error) {
	if job.jobHandle == 0 {
		return nil, errAlreadyClosed
	}
	job.handleLock.RLock()
	defer job.handleLock.RUnlock()

	info := winapi.JOBOBJECT_BASIC_ACCOUNTING_INFORMATION{}
	if err := winapi.QueryInformationJobObject(
		job.jobHandle,
		winapi.JobObjectBasicAccountingInformation,
		uintptr(unsafe.Pointer(&info)),
		uint32(unsafe.Sizeof(info)),
		nil,
	); err != nil {
		return nil, fmt.Errorf("failed to query for job object process stats: %s", err)
	}
	return &info, nil
}

// QueryStorageStats gets the storage (I/O) stats for the job object.
func (job *JobObject) QueryStorageStats() (*winapi.JOBOBJECT_BASIC_AND_IO_ACCOUNTING_INFORMATION, error) {
	if job.jobHandle == 0 {
		return nil, errAlreadyClosed
	}

	job.handleLock.RLock()
	defer job.handleLock.RUnlock()

	info := winapi.JOBOBJECT_BASIC_AND_IO_ACCOUNTING_INFORMATION{}
	if err := winapi.QueryInformationJobObject(
		job.jobHandle,
		winapi.JobObjectBasicAndIoAccountingInformation,
		uintptr(unsafe.Pointer(&info)),
		uint32(unsafe.Sizeof(info)),
		nil,
	); err != nil {
		return nil, fmt.Errorf("failed to query for job object storage stats: %s", err)
	}
	return &info, nil
}
