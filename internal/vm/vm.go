package vm

import (
	"context"
	"errors"
	"net"

	"github.com/Microsoft/hcsshim/cmd/containerd-shim-runhcs-v1/stats"
)

var (
	ErrNotSupported         = errors.New("virtstack does not support the operation")
	ErrAlreadySet           = errors.New("field has already been set")
	ErrNotInPreCreatedState = errors.New("VM is not in pre-created state")
	ErrNotInCreatedState    = errors.New("VM is not in created state")
	ErrNotInRunningState    = errors.New("VM is not in running state")
	ErrNotInPausedState     = errors.New("VM is not in paused state")
)

type UVMSource interface {
	NewLinuxUVM(ctx context.Context, id string, owner string) (UVM, error)
	NewWindowsUVM(ctx context.Context, id string, owner string) (UVM, error)
}

const (
	HCS      = "hcs"
	RemoteVM = "remotevm"
)

// UVM is an abstraction around a lightweight virtual machine. It houses core lifecycle methods such as Create
// Start, and Stop and also several optional nested interfaces that can be used to determine what the virtual machine
// supports and to configure these resources.
type UVM interface {
	// ID will return a string identifier for the Utility VM.
	ID() string

	// State returns the current running state of the Utility VM. e.g. Created, Running, Terminated
	State() State

	// Create will create the Utility VM in a paused/powered off state with whatever is present in the implementation
	// of the interfaces config at the time of the call.
	Create(ctx context.Context) error

	// Start will power on the Utility VM and put it into a running state. This will boot the guest OS and start all of the
	// devices configured on the machine.
	Start(ctx context.Context) error

	// Stop will shutdown the Utility VM and place it into a terminated state.
	Stop(ctx context.Context) error

	// Pause will place the Utility VM into a paused state. The guest OS will be halted and any devices will have be in a
	// a suspended state. Save can be used to snapshot the current state of the virtual machine, and Resume can be used to
	// place the virtual machine back into a running state.
	Pause(ctx context.Context) error

	// Resume will put a previously paused Utility VM back into a running state. The guest OS will resume operation from the point
	// in time it was paused and all devices should be un-suspended.
	Resume(ctx context.Context) error

	// Save will snapshot the state of the Utility VM at the point in time when the VM was paused.
	Save(ctx context.Context) error

	// Wait synchronously waits for the Utility VM to shutdown or terminate. A call to stop will trigger this
	// to unblock.
	Wait() error

	// Stats returns statistics about the Utility VM. This includes things like assigned memory, available memory,
	// processor runtime etc.
	Stats(ctx context.Context) (*stats.VirtualMachineStatistics, error)

	// ExitError will return any error if the Utility VM exited unexpectedly, or if the Utility VM experienced an
	// error after Wait returned, it will return the wait error.
	ExitError() error

	MemoryManager
	ProcessorManager
	BootManager
	SCSIManager
	VPMemManager
	NetworkManager
	PCIManager
	VMSocketManager
}

type State uint8

const (
	StatePreCreated State = iota
	StateCreated
	StateRunning
	StateTerminated
	StatePaused
)

type SCSIDiskType uint8

const (
	SCSIDiskTypeVHD1 SCSIDiskType = iota
	SCSIDiskTypeVHDX
	SCSIDiskTypePassThrough
)

type ProcessorManager interface {
	SetProcessorCount(ctx context.Context, count uint32) error
}

type BootManager interface {
	SetUEFIBoot(ctx context.Context, dir string, path string, args string) error
	SetLinuxKernelDirectBoot(ctx context.Context, kernel string, initRD string, cmd string) error
}

type SCSIManager interface {
	AddSCSIController(ctx context.Context, id uint32) error
	AddSCSIDisk(ctx context.Context, controller uint32, lun uint32, path string, typ SCSIDiskType, readOnly bool) error
	RemoveSCSIDisk(ctx context.Context, controller uint32, lun uint32, path string) error
}

type VPMemImageFormat uint8

const (
	VPMemImageFormatVHD1 VPMemImageFormat = iota
	VPMemImageFormatVHDX
)

type VPMemManager interface {
	AddVPMemController(ctx context.Context, maximumDevices uint32, maximumSizeBytes uint64) error
	AddVPMemDevice(ctx context.Context, id uint32, path string, readOnly bool, imageFormat VPMemImageFormat) error
	RemoveVPMemDevice(ctx context.Context, id uint32, path string) error
}

type NetworkManager interface {
	AddNIC(ctx context.Context, nicID string, endpointID string, macAddr string) error
	RemoveNIC(ctx context.Context, nicID string, endpointID string, macAddr string) error
}

// TODO dcantah: Stub for now, don't know what we need for Linux
type PCIManager interface {
	AddDevice(ctx context.Context) error
}

type VMSocketType uint8

const (
	HvSocket VMSocketType = iota
	VSock
)

type VMSocketManager interface {
	VMSocketListen(ctx context.Context, socketType VMSocketType, connID interface{}) (net.Listener, error)
}

type MemoryManager interface {
	SetMemoryLimit(ctx context.Context, memoryMB uint64) error
	SetMemoryConfig(ctx context.Context, config *MemoryConfig) error
	SetMMIOConfig(ctx context.Context, lowGapMB uint64, highBaseMB uint64, highGapMB uint64) error
}

type MemoryBackingType uint8

const (
	MemoryBackingTypeVirtual MemoryBackingType = iota
	MemoryBackingTypePhysical
)

type MemoryConfig struct {
	BackingType     MemoryBackingType
	DeferredCommit  bool
	HotHint         bool
	ColdHint        bool
	ColdDiscardHint bool
}
