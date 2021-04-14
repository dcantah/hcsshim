package hcs

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/Microsoft/hcsshim/internal/vm"
)

func (uvm *utilityVM) HVSocketListen(ctx context.Context, serviceID guid.GUID) (net.Listener, error) {
	return winio.ListenHvsock(&winio.HvsockAddr{
		VMID:      uvm.vmID,
		ServiceID: serviceID,
	})
}

func (uvm *utilityVM) VSockListen(ctx context.Context, port uint32) (net.Listener, error) {
	return nil, vm.ErrNotSupported
}
