package hcs

import (
	"context"
	"fmt"

	"github.com/Microsoft/hcsshim/internal/requesttype"
	hcsschema "github.com/Microsoft/hcsshim/internal/schema2"
	"github.com/Microsoft/hcsshim/internal/vm"
)

func (uvm *utilityVM) AddNIC(ctx context.Context, nicID, endpointID, macAddr string) error {
	if uvm.state != vm.StateRunning {
		return vm.ErrNotInRunningState
	}
	request := hcsschema.ModifySettingRequest{
		RequestType:  requesttype.Add,
		ResourcePath: fmt.Sprintf("VirtualMachine/Devices/NetworkAdapters/%s", nicID),
		Settings: hcsschema.NetworkAdapter{
			EndpointId: endpointID,
			MacAddress: macAddr,
		},
	}
	return uvm.cs.Modify(ctx, request)
}

func (uvm *utilityVM) RemoveNIC(ctx context.Context, nicID, endpointID, macAddr string) error {
	if uvm.state != vm.StateRunning {
		return vm.ErrNotInRunningState
	}
	request := hcsschema.ModifySettingRequest{
		RequestType:  requesttype.Remove,
		ResourcePath: fmt.Sprintf("VirtualMachine/Devices/NetworkAdapters/%s", nicID),
		Settings: hcsschema.NetworkAdapter{
			EndpointId: endpointID,
			MacAddress: macAddr,
		},
	}
	return uvm.cs.Modify(ctx, request)
}
