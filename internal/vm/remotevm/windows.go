package remotevm

import (
	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/Microsoft/hcsshim/internal/vm"
)

func (uvm *remoteVM) SetCPUGroup(id guid.GUID) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}

	return nil
}
