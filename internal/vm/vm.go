package vm

import (
	"context"
	"errors"
	"net"

	"github.com/Microsoft/go-winio/pkg/guid"
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

// UVMSource is an object that can launch a Windows or Linux Utility VM.
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

	WindowsConfigManager
	LinuxConfigManager
	MemoryManager
	ProcessorManager
	StorageQosManager
	SerialManager
	BootManager
	SCSIManager
	VPMemManager
	NetworkManager
	PCIManager
	VMSocketManager
}

// State refers to the states that a Utility VM can be in. The state of the Utility VM should be the source of truth for what
// operations can be performed at a given moment.
type State uint8

const (
	StatePreCreated State = iota
	StateCreated
	StateRunning
	StateTerminated
	StatePaused
)

// SCSIDiskType refers to the disk type of the scsi device. This is either a vhd, vhdx, or a physical disk.
type SCSIDiskType uint8

const (
	SCSIDiskTypeVHD1 SCSIDiskType = iota
	SCSIDiskTypeVHDX
	SCSIDiskTypePassThrough
)

// ProcessorManager handles setting and managing processor configurations for the Utility VM.
type ProcessorManager interface {
	// SetProcessorCount sets the number of virtual processors that will be assigned to the Utility VM.
	SetProcessorCount(ctx context.Context, count uint32) error
}

type MemoryBackingType uint8

const (
	MemoryBackingTypeVirtual MemoryBackingType = iota
	MemoryBackingTypePhysical
)

// MemoryConfig holds the memory options that should be configurable for a Utility VM.
type MemoryConfig struct {
	BackingType     MemoryBackingType
	DeferredCommit  bool
	HotHint         bool
	ColdHint        bool
	ColdDiscardHint bool
}

// MemoryManager handles setting and managing memory configurations for the Utility VM.
type MemoryManager interface {
	// SetMemoryLimit sets the amount of memory in megabytes that the Utility VM will be assigned.
	SetMemoryLimit(memoryMB uint64) error
	// SetMemoryConfig sets an array of different memory configuration options available. This includes things like the
	// type of memory to back the VM (virtual/physical).
	SetMemoryConfig(config *MemoryConfig) error
	// SetMMIOConfig sets memory mapped IO configurations for the Utility VM.
	SetMMIOConfig(lowGapMB uint64, highBaseMB uint64, highGapMB uint64) error
}

// BootManager manages boot configurations for the Utility VM.
type BootManager interface {
	// SetUEFIBoot sets UEFI configurations for booting a Utility VM.
	SetUEFIBoot(dir string, path string, args string) error
	// SetLinuxKernelDirectBoot sets Linux direct boot configurations for booting a Utility VM.
	SetLinuxKernelDirectBoot(kernel string, initRD string, cmd string) error
}

// SCSIManager manages adding and removing SCSI devices for a Utility VM.
type SCSIManager interface {
	// AddSCSIController adds a SCSI controller to the Utility VM configuration document.
	AddSCSIController(id uint32) error
	// AddSCSIDisk adds a SCSI disk to the configuration document if in a precreated state, or hot adds a
	// SCSI disk to the Utility VM if the VM is running.
	AddSCSIDisk(ctx context.Context, controller uint32, lun uint32, path string, typ SCSIDiskType, readOnly bool) error
	// RemoveSCSIDisk removes a SCSI disk from a Utility VM.
	RemoveSCSIDisk(ctx context.Context, controller uint32, lun uint32, path string) error
}

// VPMemImageFormat refers to the image type of the vpmem block device. This is either a vhd or vhdx.
type VPMemImageFormat uint8

const (
	VPMemImageFormatVHD1 VPMemImageFormat = iota
	VPMemImageFormatVHDX
)

// VPMemManager manages adding and removing virtual persistent memory devices for a Utility VM.
type VPMemManager interface {
	// AddVPMemController adds a new virtual pmem controller to the Utility VM.
	// `maximumDevices` specifies how many vpmem devices will be present in the guest.
	// `maximumSizeBytes` specifies the maximum size allowed for a vpmem device.
	AddVPMemController(maximumDevices uint32, maximumSizeBytes uint64) error
	// AddVPMemDevice adds a virtual pmem device to the Utility VM.
	AddVPMemDevice(ctx context.Context, id uint32, path string, readOnly bool, imageFormat VPMemImageFormat) error
	// RemoveVpmemDevice removes a virtual pmem device from the Utility VM.
	RemoveVPMemDevice(ctx context.Context, id uint32, path string) error
}

// NetworkManager manages adding and removing network adapters for a Utility VM.
type NetworkManager interface {
	// AddNIC adds a network adapter to the Utility VM. `nicID` should be a string representation of a
	// Windows GUID.
	AddNIC(ctx context.Context, nicID string, endpointID string, macAddr string) error
	// RemoveNIC removes a network adapter from the Utility VM. `nicID` should be a string representation of a
	// Windows GUID.
	RemoveNIC(ctx context.Context, nicID string, endpointID string, macAddr string) error
}

// StorageQosManager manages setting storage limits on the Utility VM.
type StorageQosManager interface {
	// SetStorageQos sets storage related options for the Utility VM
	SetStorageQos(iopsMaximum int64, bandwidthMaximum int64) error
}

// PCIManager manages assiging pci devices to a Utility VM. This is Windows specific at the moment.
type PCIManager interface {
	// AddDevice adds the pci device identified by `instanceID` to the Utility VM.
	// https://docs.microsoft.com/en-us/windows-hardware/drivers/install/instance-ids
	AddDevice(ctx context.Context, instanceID string) error
	// RemoveDevice removes the pci device identified by `instanceID` from the Utility VM.
	RemoveDevice(ctx context.Context, instanceID string) error
}

// VMSocketType refers to which hypervisor socket transport type to use.
type VMSocketType uint8

const (
	HvSocket VMSocketType = iota
	VSock
)

// VMSocketManager manages configuration for a hypervisor socket transport. This includes sockets such as
// HvSocket and Vsock.
type VMSocketManager interface {
	// VMSocketListen will create the requested vmsocket type and listen on the address specified by `connID`.
	// For HvSocket the type expected is a GUID, for Vsock it's a port of type uint32.
	VMSocketListen(ctx context.Context, socketType VMSocketType, connID interface{}) (net.Listener, error)
}

// SerialManager manages setting up serial consoles for the Utility VM.
type SerialManager interface {
	SetSerialConsole(port uint32, listenerPath string) error
}

// WindowsConfigManager manages options specific to a Windows host (cpugroups etc.)
type WindowsConfigManager interface {
	// SetCPUGroup sets the CPU group that the Utility VM will belong to on a Windows host.
	SetCPUGroup(id guid.GUID) error
}

// LinuxConfigManager manages options specific to a Linux host.
type LinuxConfigManager interface{}
