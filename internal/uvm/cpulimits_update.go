package uvm

import (
	"context"

	"github.com/Microsoft/hcsshim/internal/resourcepath"
	hcsschema "github.com/Microsoft/hcsshim/internal/schema2"
)

// UpdateCPULimits updates the CPU limits of the utility vm
func (uvm *UtilityVM) UpdateCPULimits(ctx context.Context, limits *hcsschema.ProcessorLimits) error {
	req := &hcsschema.ModifySettingRequest{
		ResourcePath: resourcepath.CPULimitsResourcePath,
		Settings:     limits,
	}

	return uvm.modify(ctx, req)
}
