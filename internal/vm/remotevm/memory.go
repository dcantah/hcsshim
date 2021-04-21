package remotevm

import (
	"github.com/Microsoft/hcsshim/internal/vm"
	"github.com/Microsoft/hcsshim/internal/vmservice"
)

func (uvm *remoteVM) SetMemoryLimit(memoryMB uint64) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}
	if uvm.config.MemoryConfig == nil {
		uvm.config.MemoryConfig = &vmservice.MemoryConfig{}
	}
	uvm.config.MemoryConfig.MemoryMb = memoryMB
	return nil
}

func (uvm *remoteVM) SetMemoryConfig(config *vm.MemoryConfig) error {
	return nil
}

func (uvm *remoteVM) SetMMIOConfig(lowGapMB uint64, highBaseMB uint64, highGapMB uint64) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}
	if uvm.config.MemoryConfig == nil {
		uvm.config.MemoryConfig = &vmservice.MemoryConfig{}
	}
	uvm.config.MemoryConfig.HighMmioBaseInMb = highBaseMB
	uvm.config.MemoryConfig.LowMmioGapInMb = lowGapMB
	uvm.config.MemoryConfig.HighMmioGapInMb = highGapMB
	return nil
}
