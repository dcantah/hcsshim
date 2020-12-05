package jobcontainers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"sync"

	"github.com/Microsoft/hcsshim/internal/cow"
	"github.com/Microsoft/hcsshim/internal/hcs"
	"github.com/Microsoft/hcsshim/internal/log"
	"golang.org/x/sys/windows"
)

// JobProcess represents a process run in a job object.
type JobProcess struct {
	cmd            *exec.Cmd
	procLock       sync.Mutex
	stdioLock      sync.Mutex
	stdin          io.WriteCloser
	stdout         io.ReadCloser
	stderr         io.ReadCloser
	waitBlock      chan struct{}
	closedWaitOnce sync.Once
	waitError      error
}

var _ cow.Process = &JobProcess{}

func newProcess(cmd *exec.Cmd) *JobProcess {
	return &JobProcess{
		cmd:       cmd,
		waitBlock: make(chan struct{}),
	}
}

func (p *JobProcess) ResizeConsole(ctx context.Context, width, height uint16) error {
	return nil
}

// Stdio returns the stdio pipes of the process
func (p *JobProcess) Stdio() (io.Writer, io.Reader, io.Reader) {
	return p.stdin, p.stdout, p.stderr
}

// Signal sends a signal to the process.
func (p *JobProcess) Signal(ctx context.Context, options interface{}) (bool, error) {
	p.procLock.Lock()
	defer p.procLock.Unlock()

	if p.exited() {
		return false, errors.New("signal not sent. process already exited")
	}

	if err := windows.GenerateConsoleCtrlEvent(windows.CTRL_BREAK_EVENT, uint32(p.Pid())); err != nil {
		return false, fmt.Errorf("failed to send signal: %s", err)
	}
	return true, nil
}

// CloseStdin closes the stdin pipe of the process.
func (p *JobProcess) CloseStdin(ctx context.Context) error {
	p.stdioLock.Lock()
	defer p.stdioLock.Unlock()
	return p.stdin.Close()
}

// Wait waits for the process to exit. If the process has already exited returns
// the previous error (if any).
func (p *JobProcess) Wait() error {
	<-p.waitBlock
	return p.waitError
}

// Start starts the job object process
func (p *JobProcess) Start() error {
	return p.cmd.Start()
}

// This should only be called once.
func (p *JobProcess) waitBackground(ctx context.Context) {
	log.G(ctx).WithField("pid", p.Pid()).Debug("waitBackground for JobProcess")

	var err error
	if p.cmd.Process == nil {
		err = errors.New("process has not been started")
	}

	// Wait for process to get signaled/exit/terminate/.
	err = p.cmd.Wait()

	// Wait closes the stdio pipes so theres no need to later on.
	p.stdioLock.Lock()
	p.stdin = nil
	p.stdout = nil
	p.stderr = nil
	p.stdioLock.Unlock()

	p.closedWaitOnce.Do(func() {
		p.waitError = err
		close(p.waitBlock)
	})
}

// ExitCode returns the exit code of the process.
func (p *JobProcess) ExitCode() (int, error) {
	p.procLock.Lock()
	defer p.procLock.Unlock()

	if !p.exited() {
		return -1, errors.New("process has not exited")
	}
	return p.cmd.ProcessState.ExitCode(), nil
}

// Pid returns the processes PID
func (p *JobProcess) Pid() int {
	if process := p.cmd.Process; process != nil {
		return process.Pid
	}
	return 0
}

// Close cleans up any state associated with the process but does not kill it.
func (p *JobProcess) Close() error {
	p.stdioLock.Lock()
	if p.stdin != nil {
		p.stdin.Close()
		p.stdin = nil
	}
	if p.stdout != nil {
		p.stdout.Close()
		p.stdout = nil
	}
	if p.stderr != nil {
		p.stderr.Close()
		p.stderr = nil
	}
	p.stdioLock.Unlock()

	p.closedWaitOnce.Do(func() {
		p.waitError = hcs.ErrAlreadyClosed
		close(p.waitBlock)
	})
	return nil
}

// Kill kills the running process. Go calls TerminateProcess under the hood.
func (p *JobProcess) Kill(ctx context.Context) (bool, error) {
	log.G(ctx).WithField("pid", p.Pid()).Debug("killing job process")

	p.procLock.Lock()
	defer p.procLock.Unlock()

	if p.exited() {
		return false, errors.New("kill not sent. process already exited")
	}

	if p.cmd.Process != nil {
		if err := p.cmd.Process.Kill(); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (p *JobProcess) exited() bool {
	if p.cmd.ProcessState == nil {
		return false
	}
	return p.cmd.ProcessState.Exited()
}
