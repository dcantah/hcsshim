package remotevm

import (
	"context"
	"io"
	"net"
	"os"
	"os/exec"

	"github.com/Microsoft/hcsshim/internal/jobobject"
	"github.com/Microsoft/hcsshim/internal/log"
	"github.com/Microsoft/hcsshim/internal/vm"
	"github.com/Microsoft/hcsshim/internal/vmservice"
	"github.com/containerd/ttrpc"
	ptypes "github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type source struct {
	// Path to the binary to launch
	binary string
	// Address of the server process implementing the vmservice interface.
	addr string
}

func NewSource(binary, addr string) (vm.UVMSource, error) {
	return &source{
		binary: binary,
		addr:   addr,
	}, nil
}

func (s *source) NewLinuxUVM(ctx context.Context, id, owner string) (vm.UVM, error) {
	var job *jobobject.JobObject
	if s.binary != "" {
		log.G(ctx).WithFields(logrus.Fields{
			"binary":  s.binary,
			"address": s.addr,
		}).Debug("starting remotevm server process")

		opts := &jobobject.Options{
			Name: id,
		}
		job, err := jobobject.Create(ctx, opts)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create job object for remotevm process")
		}

		// This expects the remotevm server process to have a ttrpc flag. Possibly revisit this.s
		cmd := exec.Command(s.binary, "--ttrpc", s.addr)
		p, err := cmd.StdoutPipe()
		if err != nil {
			return nil, errors.Wrap(err, "failed to create stdout pipe")
		}

		if err := cmd.Start(); err != nil {
			return nil, errors.Wrap(err, "failed to start remotevm server process")
		}

		if err := job.Assign(uint32(cmd.Process.Pid)); err != nil {
			return nil, errors.Wrap(err, "failed to assign remotevm process to job")
		}

		if err := job.SetTerminateOnLastHandleClose(); err != nil {
			return nil, errors.Wrap(err, "failed to set terminate on last handle closed for remotevm job object")
		}

		f, err := os.Open(os.DevNull)
		if err != nil {
			return nil, errors.Wrap(err, "failed to open nul file")
		}

		// Wait for stdout to close. This is our signal that the server is successfully up and running.
		_, _ = io.Copy(f, p)
	}

	conn, err := net.Dial("unix", s.addr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial remotevm address %q", s.addr)
	}

	c := ttrpc.NewClient(conn, ttrpc.WithOnClose(func() { conn.Close() }))
	vmClient := vmservice.NewVMClient(c)

	return &remoteVM{
		id: id,
		config: &vmservice.VMConfig{
			MemoryConfig:    &vmservice.MemoryConfig{},
			DevicesConfig:   &vmservice.DevicesConfig{},
			ProcessorConfig: &vmservice.ProcessorConfig{},
			SerialConfig:    &vmservice.SerialConfig{},
			ExtraData:       make(map[string]string),
		},
		job:    job,
		client: vmClient,
	}, nil
}

func (s *source) NewWindowsUVM(ctx context.Context, id, owner string) (vm.UVM, error) {
	return nil, vm.ErrNotSupported
}

type remoteVM struct {
	id        string
	state     vm.State
	waitError error
	job       *jobobject.JobObject
	config    *vmservice.VMConfig
	client    vmservice.VMService
}

func (uvm *remoteVM) ID() string {
	return uvm.id
}

func (uvm *remoteVM) State() vm.State {
	return uvm.state
}

func (uvm *remoteVM) Create(ctx context.Context) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}
	if _, err := uvm.client.CreateVM(ctx, &vmservice.CreateVMRequest{Config: uvm.config, LogID: uvm.ID()}); err != nil {
		return errors.Wrap(err, "failed to create remote VM")
	}
	uvm.state = vm.StateCreated
	return nil
}

func (uvm *remoteVM) Start(ctx context.Context) error {
	if uvm.state != vm.StateCreated {
		return vm.ErrNotInCreatedState
	}
	// The VM is expected to be in a pause state after Create, so start is truthfully just resuming the
	// VM. This is really what HCS does behind the scenes anyways, it's just labeled as Start.
	if _, err := uvm.client.ResumeVM(ctx, &ptypes.Empty{}); err != nil {
		return errors.Wrap(err, "failed to start remote VM")
	}
	uvm.state = vm.StateRunning
	return nil
}

func (uvm *remoteVM) Stop(ctx context.Context) error {
	if uvm.state != vm.StateRunning {
		return vm.ErrNotInRunningState
	}
	if _, err := uvm.client.TeardownVM(ctx, &ptypes.Empty{}); err != nil {
		return errors.Wrap(err, "failed to stop remote VM")
	}
	uvm.state = vm.StateTerminated
	return nil
}

func (uvm *remoteVM) Wait() error {
	if uvm.state != vm.StateRunning {
		return vm.ErrNotInRunningState
	}
	if uvm.state == vm.StateTerminated {
		return nil
	}
	_, err := uvm.client.WaitVM(context.Background(), &ptypes.Empty{})
	if err != nil {
		uvm.waitError = err
		return errors.Wrap(err, "failed to wait on remote VM")
	}
	return nil
}

func (uvm *remoteVM) Pause(ctx context.Context) error {
	return vm.ErrNotSupported
}

func (uvm *remoteVM) Resume(ctx context.Context) error {
	if uvm.state != vm.StatePaused {
		return vm.ErrNotInPausedState
	}
	// Unlike HCS, resume can be called both after create to power on the devices/boot the guest OS
	// and also after pausing the VM.
	if _, err := uvm.client.ResumeVM(ctx, &ptypes.Empty{}); err != nil {
		return errors.Wrap(err, "failed to resume remote VM")
	}
	uvm.state = vm.StateRunning
	return nil
}

func (uvm *remoteVM) Save(ctx context.Context) error {
	return vm.ErrNotSupported
}

func (uvm *remoteVM) ExitError() error {
	return uvm.waitError
}
