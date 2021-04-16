package remotevm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/Microsoft/hcsshim/hcn"
	"github.com/Microsoft/hcsshim/internal/vm"
	"github.com/Microsoft/hcsshim/internal/vmservice"
	"github.com/pkg/errors"
)

func (uvm *remoteVM) AddNIC(ctx context.Context, nicID, endpointID, macAddr string) error {
	portID, err := guid.NewV4()
	if err != nil {
		return errors.Wrap(err, "failed to generate guid for port")
	}

	vmEndpointRequest := hcn.VmEndpointRequest{
		PortId:           portID,
		VirtualNicName:   fmt.Sprintf("%s--%s", nicID, portID),
		VirtualMachineId: guid.GUID{},
	}

	m, err := json.Marshal(vmEndpointRequest)
	if err != nil {
		return errors.Wrap(err, "failed to marshal endpoint request json")
	}

	if err := hcn.ModifyEndpointSettings(endpointID, &hcn.ModifyEndpointSettingRequest{
		ResourceType: hcn.EndpointResourceTypePort,
		RequestType:  hcn.RequestTypeAdd,
		Settings:     json.RawMessage(m),
	}); err != nil {
		return errors.Wrap(err, "failed to configure switch port")
	}

	// Get updated endpoint with new fields (need switch ID)
	ep, err := hcn.GetEndpointByID(endpointID)
	if err != nil {
		return errors.Wrapf(err, "failed to get endpoint %q", endpointID)
	}

	type ExtraInfo struct {
		Allocators []struct {
			SwitchId         string
			EndpointPortGuid string
		}
	}

	var exi ExtraInfo
	if err := json.Unmarshal(ep.Health.Extra.Resources, &exi); err != nil {
		return errors.Wrapf(err, "failed to unmarshal resource data from endpoint %q", endpointID)
	}

	if len(exi.Allocators) != 1 {
		return errors.New("no resource data found for endpoint")
	}

	nic := &vmservice.NICConfig{
		MacAddress: macAddr,
		PortID:     portID.String(),
		SwitchID:   exi.Allocators[0].SwitchId,
	}

	switch uvm.state {
	case vm.StatePreCreated:
		uvm.config.DevicesConfig.NicConfig = append(uvm.config.DevicesConfig.NicConfig, nic)
	case vm.StateRunning:
		// Hot add
		if _, err := uvm.client.ModifyResource(ctx,
			&vmservice.ModifyResourceRequest{
				Type: vmservice.ModifyType_ADD,
				Resource: &vmservice.ModifyResourceRequest_NicConfig{
					NicConfig: nic,
				},
			},
		); err != nil {
			return errors.Wrap(err, "failed to add network adapter")
		}
	default:
		return errors.New("VM is not in pre-created or running state")
	}

	return nil
}
