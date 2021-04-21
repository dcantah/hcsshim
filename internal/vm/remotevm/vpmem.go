package remotevm

import (
	"context"

	"github.com/Microsoft/hcsshim/internal/vm"
)

func (uvm *remoteVM) AddVPMemController(maximumDevices uint32, maximumSizeBytes uint64) error {
	return vm.ErrNotSupported
}

func (uvm *remoteVM) AddVPMemDevice(ctx context.Context, id uint32, path string, readOnly bool, imageFormat vm.VPMemImageFormat) error {
	return vm.ErrNotSupported
}

func (uvm *remoteVM) RemoveVPMemDevice(ctx context.Context, id uint32, path string) error {
	return vm.ErrNotSupported
}
