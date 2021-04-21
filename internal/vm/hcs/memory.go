package hcs

import (
	"github.com/Microsoft/hcsshim/internal/vm"
)

func (uvm *utilityVM) SetMemoryLimit(memoryMB uint64) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}
	uvm.doc.VirtualMachine.ComputeTopology.Memory.SizeInMB = memoryMB
	return nil
}

func (uvm *utilityVM) SetMemoryConfig(config *vm.MemoryConfig) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}
	memory := uvm.doc.VirtualMachine.ComputeTopology.Memory
	memory.AllowOvercommit = config.BackingType == vm.MemoryBackingTypeVirtual
	memory.EnableDeferredCommit = config.DeferredCommit
	memory.EnableHotHint = config.HotHint
	memory.EnableColdHint = config.ColdHint
	memory.EnableColdDiscardHint = config.ColdDiscardHint
	return nil
}

func (uvm *utilityVM) SetMMIOConfig(lowGapMB uint64, highBaseMB uint64, highGapMB uint64) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}
	memory := uvm.doc.VirtualMachine.ComputeTopology.Memory
	memory.LowMMIOGapInMB = lowGapMB
	memory.HighMMIOBaseInMB = highBaseMB
	memory.HighMMIOGapInMB = highGapMB
	return nil
}
