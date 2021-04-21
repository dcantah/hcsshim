package remotevm

import (
	"github.com/Microsoft/hcsshim/internal/vm"
	"github.com/Microsoft/hcsshim/internal/vmservice"
)

func (uvm *remoteVM) SetLinuxKernelDirectBoot(kernel string, initRD string, cmd string) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}
	if uvm.config.BootConfig != nil {
		return vm.ErrAlreadySet
	}
	uvm.config.BootConfig = &vmservice.VMConfig_DirectBoot{
		DirectBoot: &vmservice.DirectBoot{
			KernelPath:    kernel,
			InitrdPath:    initRD,
			KernelCmdline: cmd,
		},
	}
	return nil
}

func (uvm *remoteVM) SetUEFIBoot(dir string, path string, args string) error {
	return vm.ErrNotSupported
}
