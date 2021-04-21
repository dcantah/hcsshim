package hcs

import (
	"github.com/Microsoft/hcsshim/internal/vm"
)

func (uvm *utilityVM) SetStorageQos(iopsMaximum int64, bandwidthMaximum int64) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}

	uvm.doc.VirtualMachine.StorageQoS.BandwidthMaximum = int32(bandwidthMaximum)
	uvm.doc.VirtualMachine.StorageQoS.IopsMaximum = int32(iopsMaximum)
	return nil
}
