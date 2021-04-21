package remotevm

import (
	"context"

	"github.com/Microsoft/hcsshim/internal/vm"
)

func (uvm *remoteVM) AddDevice(ctx context.Context, instanceID string) error {
	return vm.ErrNotSupported
}

func (uvm *remoteVM) RemoveDevice(ctx context.Context, instanceID string) error {
	return vm.ErrNotSupported
}
