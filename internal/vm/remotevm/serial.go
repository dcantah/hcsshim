package remotevm

import (
	"github.com/Microsoft/hcsshim/internal/vm"
	"github.com/Microsoft/hcsshim/internal/vmservice"
)

func (uvm *remoteVM) SetSerialConsole(port uint32, listenerPath string) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}

	config := &vmservice.SerialConfig_Config{
		Port:       port,
		SocketPath: listenerPath,
	}
	uvm.config.SerialConfig.Ports = []*vmservice.SerialConfig_Config{config}
	return nil
}
