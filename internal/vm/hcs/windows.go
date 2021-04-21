package hcs

import (
	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/Microsoft/hcsshim/internal/vm"
)

func (uvm *utilityVM) SetCPUGroup(id guid.GUID) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}

	return nil
}
