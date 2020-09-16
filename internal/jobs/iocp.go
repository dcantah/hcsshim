package jobs

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/Microsoft/hcsshim/internal/winapi"
	"golang.org/x/sys/windows"
)

var (
	ioInitOnce       sync.Once
	ioCompletionPort windows.Handle
	// Map of job object handle values to the channel it will receive notifications
	// on. If job object creation succeeds this will get a value added to it.
	jobMap = make(map[uint32]chan IOCPNotif)
)

type IOCPNotif struct {
	Code uint32
	Err  error
}

// ioCompletionProcessor processes completed async IOs forever
func pollIOCP(iocpHandle windows.Handle) {
	// Code will hold the value of the job object specific message that was received.
	// Key will hold the job object handle value so we can determine what job the
	// notification is for.
	var (
		overlapped uintptr
		code       uint32
		key        uint32
	)

	for {
		// Poll IOCP and dispatch messages to job object the notification belongs to.
		if err := windows.GetQueuedCompletionStatus(iocpHandle, &code, &key, (**windows.Overlapped)(unsafe.Pointer(&overlapped)), windows.INFINITE); err == nil {
			fmt.Println("Code received: ", code, " Key: ", key)
			if jobChan, ok := jobMap[key]; ok {
				jobChan <- IOCPNotif{
					Code: code,
					Err:  err,
				}
			}
		}
	}
}

func initIo() {
	h, err := windows.CreateIoCompletionPort(windows.InvalidHandle, 0, 0, 0xffffffff)
	// If this fails we don't have anyway of monitoring when a job container has
	// shutdown. This is run once via the sync.Once.Do method so we can't just
	// return an error here.
	if err != nil {
		panic(err)
	}
	ioCompletionPort = h
	go pollIOCP(h)
}

// Assigns an IO completion port to get notified of events for the registered job
// object.
func attachIOCP(job windows.Handle, iocp windows.Handle) (int, error) {
	info := winapi.JOBOBJECT_ASSOCIATE_COMPLETION_PORT{
		CompletionKey:  job,
		CompletionPort: iocp,
	}
	return windows.SetInformationJobObject(job, windows.JobObjectAssociateCompletionPortInformation, uintptr(unsafe.Pointer(&info)), uint32(unsafe.Sizeof(info)))
}
