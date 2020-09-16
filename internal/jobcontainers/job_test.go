package jobcontainers

import (
	"context"
	"fmt"
	"os/exec"
	"syscall"
	"testing"
	"time"

	"github.com/Microsoft/hcsshim/internal/jobs"
	"github.com/Microsoft/hcsshim/internal/winapi"
	"golang.org/x/sys/windows"
)

// Helper to create processes to assign to job object tests
func createProcesses(count int) []*JobProcess {
	procs := make([]*JobProcess, count)
	for i := 0; i < count; i++ {
		// cmd := exec.Command("ping", "-t", "127.0.0.1")
		// cmd := exec.Command("winver")
		cmd := exec.Command(`C:\Users\dcanter\go\src\github.com\dcantah\demo\demo.exe`)
		cmd.SysProcAttr = &syscall.SysProcAttr{
			CreationFlags: windows.CREATE_NEW_PROCESS_GROUP,
		}
		procs[i] = newProcess(cmd)
	}
	return procs
}

func startAndWait(job *jobs.JobObject, procs []*JobProcess) error {
	for _, proc := range procs {
		if err := proc.Start(); err != nil {
			return err
		}
		if err := job.Assign(uint32(proc.Pid())); err != nil {
			return err
		}
		go proc.waitBackground(context.Background())
	}
	return nil
}

func cleanup(job *jobs.JobObject, procs []*JobProcess) {
	job.Close()
	for _, proc := range procs {
		if code, _ := proc.ExitCode(); code == -1 {
			proc.Kill(context.Background())
		}
	}
}

// func TestCreateAndTerminateJob(t *testing.T) {
// 	job, err := jobs.CreateJobObject("test")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs := createProcesses(2)

// 	defer func() {
// 		cleanup(job, procs)
// 	}()

// 	if err := startAndWait(job, procs); err != nil {
// 		t.Fatalf("failed to start and wait on processes")
// 	}

// 	if err := job.Terminate(); err != nil {
// 		t.Fatalf("failed to terminate job object: %s", err)
// 	}
// }

// func TestCreateAndShutdownJob(t *testing.T) {
// 	job, err := jobs.CreateJobObject("test")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs := createProcesses(2)

// 	defer func() {
// 		cleanup(job, procs)
// 	}()

// 	if err := startAndWait(job, procs); err != nil {
// 		t.Fatalf("failed to start and wait on processes")
// 	}

// 	if err := job.Shutdown(context.Background()); err != nil {
// 		t.Fatalf("failed to shutdown job object: %s", err)
// 	}
// }

// func TestJobPIDs(t *testing.T) {
// 	job, err := jobs.CreateJobObject("test")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs := createProcesses(2)

// 	defer func() {
// 		cleanup(job, procs)
// 	}()

// 	if err := startAndWait(job, procs); err != nil {
// 		t.Fatalf("failed to start and wait on processes")
// 	}

// 	pidsMap := make(map[int]struct{})
// 	for _, proc := range procs {
// 		pidsMap[proc.Pid()] = struct{}{}
// 	}

// 	pids, err := job.Pids()
// 	if err != nil {
// 		t.Fatalf("failed to get PIDs in job: %s", err)
// 	}

// 	if len(pids) != len(procs) {
// 		t.Fatalf("number of PIDs in job incorrect")
// 	}

// 	for i := 0; i < len(pids); i++ {
// 		if _, ok := pidsMap[int(pids[i])]; !ok {
// 			t.Fatalf("PID not present in job object")
// 		}
// 	}

// 	if err := job.Shutdown(context.Background()); err != nil {
// 		t.Fatalf("failed to shutdown job: %s", err)
// 	}
// }

// func TestJobIOCP(t *testing.T) {
// 	job, err := jobs.CreateJobObject("test")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs := createProcesses(2)

// 	defer func() {
// 		cleanup(job, procs)
// 	}()

// 	if err := startAndWait(job, procs); err != nil {
// 		t.Fatalf("failed to start and wait on processes")
// 	}

// 	iocpChan := job.IOCPChan()

// 	for {
// 		notif := <-iocpChan
// 		if notif.Err != nil {
// 			t.Fatalf("failed to poll IOCP: %s", err)
// 		}
// 		if notif.Code == winapi.JOB_OBJECT_MSG_ACTIVE_PROCESS_ZERO {
// 			break
// 		}
// 	}
// }

// func TestMultipleJobsIOCP(t *testing.T) {
// 	job, err := jobs.CreateJobObject("test")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs := createProcesses(2)

// 	job2, err := jobs.CreateJobObject("test2")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs2 := createProcesses(2)

// 	defer func() {
// 		cleanup(job, procs)
// 		cleanup(job2, procs2)
// 	}()

// 	if err := startAndWait(job, procs); err != nil {
// 		t.Fatalf("failed to start and wait on processes")
// 	}

// 	iocpChan := job.IOCPChan()
// 	iocpChan2 := job2.IOCPChan()

// 	notif <- iocpChan
// 	notif2 <- iocpChan2
// 	fmt.Println()
// }

// func TestJobNotificationShutdown(t *testing.T) {
// 	job, err := jobs.CreateJobObject("test")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs := createProcesses(2)

// 	defer func() {
// 		cleanup(job, procs)
// 	}()

// 	if err := startAndWait(job, procs); err != nil {
// 		t.Fatalf("failed to start and wait on processes")
// 	}

// 	done := make(chan struct{})
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	go func() {
// 		// Wait two seconds and then kill all the procs. Should trigger the IOCP
// 		// message if all the processes exit successfully
// 		time.Sleep(time.Second * 2)
// 		if err := job.Shutdown(context.Background()); err != nil {
// 			t.Fatalf("failed to shutdown job: %s", err)
// 		}
// 	}()

// 	go func() {
// 		for {
// 			code, err := job.PollIOCP()
// 			if err != nil {
// 				t.Fatalf("failed to poll IOCP: %s", err)
// 			}

// 			// No more processes in the job. Success state
// 			if code == winapi.JOB_OBJECT_MSG_ACTIVE_PROCESS_ZERO {
// 				close(done)
// 				return
// 			}
// 		}
// 	}()

// 	// If we still haven't received the all processes exited message after context
// 	// timeout fail the test.
// 	select {
// 	case <-done:
// 		break
// 	case <-ctx.Done():
// 		t.Fatal(ctx.Err())
// 	}
// }

// func TestJobNotificationKill(t *testing.T) {
// 	job, err := jobs.CreateJobObject("test")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs := createProcesses(2)

// 	defer func() {
// 		cleanup(job, procs)
// 	}()

// 	if err := startAndWait(job, procs); err != nil {
// 		t.Fatalf("failed to start and wait on processes")
// 	}

// 	done := make(chan struct{})
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	go func() {
// 		// Wait two seconds and then kill all the procs. Should trigger the IOCP
// 		// message if all the processes exit successfully
// 		time.Sleep(time.Second * 2)
// 		for _, proc := range procs {
// 			proc.Kill(context.Background())
// 		}
// 	}()

// 	go func() {
// 		for {
// 			code, err := job.PollIOCP()
// 			if err != nil {
// 				t.Fatalf("failed to poll IOCP: %s", err)
// 			}

// 			// No more processes in the job. Success state
// 			if code == winapi.JOB_OBJECT_MSG_ACTIVE_PROCESS_ZERO {
// 				close(done)
// 				return
// 			}
// 		}
// 	}()

// 	// If we still haven't received the all processes exited message after context
// 	// timeout fail the test.
// 	select {
// 	case <-done:
// 		break
// 	case <-ctx.Done():
// 		t.Fatal(ctx.Err())
// 	}
// }

// func TestKillProcess(t *testing.T) {
// 	job, err := jobs.CreateJobObject("test")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs := createProcesses(1)

// 	if err := startAndWait(job, procs); err != nil {
// 		t.Fatalf("failed to start and wait on processes")
// 	}

// 	defer func() {
// 		job.Terminate()
// 		job.Close()
// 	}()

// 	for _, proc := range procs {
// 		if code, _ := proc.ExitCode(); code == -1 {
// 			if _, err := proc.Kill(context.Background()); err != nil {
// 				t.Fatalf("failed to kill process: %s", err)
// 			}
// 		}
// 	}
// }

// func TestJobStats(t *testing.T) {
// 	job, err := jobs.CreateJobObject("test")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs := createProcesses(2)

// 	defer func() {
// 		cleanup(job, procs)
// 	}()

// 	if err := startAndWait(job, procs); err != nil {
// 		t.Fatalf("failed to start and wait on processes")
// 	}

// 	_, err = job.QueryMemoryStats()
// 	if err != nil {
// 		t.Fatalf("failed to query for memory stats on job object: %s", err)
// 	}

// 	_, err = job.QueryProcessorStats()
// 	if err != nil {
// 		t.Fatalf("failed to query for processor stats on job object: %s", err)
// 	}

// 	_, err = job.QueryStorageStats()
// 	if err != nil {
// 		t.Fatalf("failed to query for storage stats on job object: %s", err)
// 	}

// 	if err := job.Shutdown(context.Background()); err != nil {
// 		t.Fatalf("failed to shutdown job object: %s", err)
// 	}
// }

// func TestJobIOCP(t *testing.T) {
// 	job, err := jobs.CreateJobObject("test")
// 	if err != nil {
// 		t.Fatalf("failed to create job object: %s", err)
// 	}

// 	procs := createProcesses(2)

// 	defer func() {
// 		cleanup(job, procs)
// 	}()

// 	if err := startAndWait(job, procs); err != nil {
// 		t.Fatalf("failed to start and wait on processes")
// 	}

// 	iocpChan := job.IOCPChan()

// 	go func() {
// 		// Wait two seconds and then kill all the procs. Should trigger the IOCP
// 		// message if all the processes exit successfully
// 		time.Sleep(time.Second * 2)
// 		if err := job.Shutdown(context.Background()); err != nil {
// 			fmt.Println("Did something go wrong?")
// 			t.Fatalf("failed to shutdown job: %s", err)
// 		}
// 	}()

// 	for {
// 		notif := <-iocpChan
// 		fmt.Println("notif received: ", notif)
// 		if notif.Err != nil {
// 			t.Fatalf("failed to poll IOCP: %s", err)
// 		}
// 		if notif.Code == winapi.JOB_OBJECT_MSG_ACTIVE_PROCESS_ZERO {
// 			break
// 		}
// 	}
// }

func TestMultipleJobsIOCP(t *testing.T) {
	job, err := jobs.CreateJobObject("test")
	if err != nil {
		t.Fatalf("failed to create job object: %s", err)
	}

	procs := createProcesses(3)

	defer func() {
		cleanup(job, procs)
	}()

	if err := startAndWait(job, procs); err != nil {
		t.Fatalf("failed to start and wait on processes")
	}

	iocpChan := job.IOCPChan()

	job2, err := jobs.CreateJobObject("test2")
	if err != nil {
		t.Fatalf("failed to create job object: %s", err)
	}

	defer job2.Close()

	procs2 := createProcesses(4)

	defer func() {
		cleanup(job2, procs2)
	}()

	if err := startAndWait(job2, procs2); err != nil {
		t.Fatalf("failed to start and wait on processes")
	}

	iocpChan2 := job2.IOCPChan()

	go func() {
		// Wait two seconds and then kill all the procs. Should trigger the IOCP
		// message if all the processes exit successfully
		time.Sleep(time.Second * 2)
		if err := job.Shutdown(context.Background()); err != nil {
			t.Fatalf("failed to shutdown job: %s", err)
		}
		if err := job2.Shutdown(context.Background()); err != nil {
			t.Fatalf("failed to shutdown job: %s", err)
		}
	}()

	for {
		notif := <-iocpChan
		fmt.Println("notif received in job 1: ", notif)
		if notif.Err != nil {
			t.Fatalf("failed to poll IOCP: %s", err)
		}
		if notif.Code == winapi.JOB_OBJECT_MSG_ACTIVE_PROCESS_ZERO {
			fmt.Println("Im out bro")
			break
		}
	}

	for {
		notif := <-iocpChan2
		fmt.Println("notif received in job 2: ", notif)
		if notif.Err != nil {
			t.Fatalf("failed to poll IOCP: %s", err)
		}
		if notif.Code == winapi.JOB_OBJECT_MSG_ACTIVE_PROCESS_ZERO {
			break
		}
	}
}
