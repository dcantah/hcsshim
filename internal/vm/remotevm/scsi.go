package remotevm

import (
	"context"
	"fmt"

	"github.com/Microsoft/hcsshim/internal/vm"
	"github.com/Microsoft/hcsshim/internal/vmservice"
	"github.com/pkg/errors"
)

func (uvm *remoteVM) AddSCSIController(ctx context.Context, id uint32) error {
	// TODO: dcantah
	return nil
}

func (uvm *remoteVM) AddSCSIDisk(ctx context.Context, controller uint32, lun uint32, path string, typ vm.SCSIDiskType, readOnly bool) error {
	var diskType vmservice.DiskType
	switch typ {
	case vm.SCSIDiskTypeVHD1:
		diskType = vmservice.DiskType_SCSI_DISK_TYPE_VHD1
	case vm.SCSIDiskTypeVHDX:
		diskType = vmservice.DiskType_SCSI_DISK_TYPE_VHDX
	case vm.SCSIDiskTypePassThrough:
		diskType = vmservice.DiskType_SCSI_DISK_TYPE_PHYSICAL
	default:
		return fmt.Errorf("unsupported SCSI disk type: %d", typ)
	}

	disk := &vmservice.SCSIDisk{
		Controller: controller,
		Lun:        lun,
		HostPath:   path,
		Type:       diskType,
		ReadOnly:   readOnly,
	}

	switch uvm.state {
	case vm.StatePreCreated:
		uvm.config.DevicesConfig.ScsiDisks = append(uvm.config.DevicesConfig.ScsiDisks, disk)
	case vm.StateRunning:
		if _, err := uvm.client.ModifyResource(ctx,
			&vmservice.ModifyResourceRequest{
				Type: vmservice.ModifyType_ADD,
				Resource: &vmservice.ModifyResourceRequest_ScsiDisk{
					ScsiDisk: disk,
				},
			},
		); err != nil {
			return errors.Wrap(err, "failed to add SCSI disk")
		}
	default:
		return errors.New("VM is not in pre-created or running state")
	}

	return nil
}

func (uvm *remoteVM) RemoveSCSIDisk(ctx context.Context, controller uint32, lun uint32, path string) error {
	if uvm.state != vm.StateRunning {
		return vm.ErrNotInRunningState
	}

	disk := &vmservice.SCSIDisk{
		Controller: controller,
		Lun:        lun,
		HostPath:   path,
	}

	if _, err := uvm.client.ModifyResource(ctx,
		&vmservice.ModifyResourceRequest{
			Type: vmservice.ModifyType_REMOVE,
			Resource: &vmservice.ModifyResourceRequest_ScsiDisk{
				ScsiDisk: disk,
			},
		},
	); err != nil {
		return errors.Wrapf(err, "failed to remove SCSI disk %q", path)
	}

	return nil
}
