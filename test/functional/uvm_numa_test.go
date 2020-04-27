// +build functional uvmproperties

package functional

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	hcsschema "github.com/Microsoft/hcsshim/internal/schema2"
	"github.com/Microsoft/hcsshim/internal/uvm"
	testutilities "github.com/Microsoft/hcsshim/test/functional/utilities"
)

func TestNUMANodeCount_LCOW(t *testing.T) {
	testutilities.RequiresBuild(t, 18943)

	numa := &hcsschema.Numa{
		VirtualNodeCount: 2,
	}
	data, err := json.Marshal(numa)
	if err != nil {
		t.Fatalf("failed to marshal numa settings: %s", err)
	}
	opts := uvm.NewDefaultOptionsLCOW(t.Name(), "")
	opts.NUMATopologyJSON = string(data)
	uvm := testutilities.CreateLCOWUVMFromOpts(context.Background(), t, opts)
	defer uvm.Close()

	stats, err := uvm.Stats(context.Background())
	if err != nil {
		t.Fatalf("failed to retrieve UVM memory stats: %s", err)
	}
	if stats.Memory.VirtualNodeCount != uint32(numa.VirtualNodeCount) {
		t.Fatalf("virtual node count incorrect. expected: %d but got %d", numa.VirtualNodeCount, stats.Memory.VirtualNodeCount)
	}
}

func TestNUMANodeCount_WCOW_Hypervisor(t *testing.T) {
	testutilities.RequiresBuild(t, 18943)

	numa := &hcsschema.Numa{
		VirtualNodeCount: 2,
	}
	data, err := json.Marshal(numa)
	if err != nil {
		t.Fatalf("failed to marshal numa settings: %s", err)
	}
	opts := uvm.NewDefaultOptionsWCOW(t.Name(), "")
	opts.NUMATopologyJSON = string(data)
	uvm, _, uvmScratchDir := testutilities.CreateWCOWUVMCustom(context.Background(), t, t.Name(), "microsoft/nanoserver", opts)
	defer os.RemoveAll(uvmScratchDir)
	defer uvm.Close()

	stats, err := uvm.Stats(context.Background())
	if err != nil {
		t.Fatalf("failed to retrieve UVM memory stats: %s", err)
	}
	if stats.Memory.VirtualNodeCount != uint32(numa.VirtualNodeCount) {
		t.Fatalf("virtual node count incorrect. expected: %d but got %d", numa.VirtualNodeCount, stats.Memory.VirtualNodeCount)
	}
}
