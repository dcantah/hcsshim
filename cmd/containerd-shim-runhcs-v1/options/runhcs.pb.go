// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/Microsoft/hcsshim/cmd/containerd-shim-runhcs-v1/options/runhcs.proto

package options

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Options_DebugType int32

const (
	Options_NPIPE Options_DebugType = 0
	Options_FILE  Options_DebugType = 1
	Options_ETW   Options_DebugType = 2
)

var Options_DebugType_name = map[int32]string{
	0: "NPIPE",
	1: "FILE",
	2: "ETW",
}

var Options_DebugType_value = map[string]int32{
	"NPIPE": 0,
	"FILE":  1,
	"ETW":   2,
}

func (x Options_DebugType) String() string {
	return proto.EnumName(Options_DebugType_name, int32(x))
}

func (Options_DebugType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b643df6839c75082, []int{0, 0}
}

type Options_SandboxIsolation int32

const (
	Options_PROCESS    Options_SandboxIsolation = 0
	Options_HYPERVISOR Options_SandboxIsolation = 1
)

var Options_SandboxIsolation_name = map[int32]string{
	0: "PROCESS",
	1: "HYPERVISOR",
}

var Options_SandboxIsolation_value = map[string]int32{
	"PROCESS":    0,
	"HYPERVISOR": 1,
}

func (x Options_SandboxIsolation) String() string {
	return proto.EnumName(Options_SandboxIsolation_name, int32(x))
}

func (Options_SandboxIsolation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b643df6839c75082, []int{0, 1}
}

// Options are the set of customizations that can be passed at Create time.
type Options struct {
	// enable debug tracing
	Debug bool `protobuf:"varint,1,opt,name=debug,proto3" json:"debug,omitempty"`
	// debug tracing output type
	DebugType Options_DebugType `protobuf:"varint,2,opt,name=debug_type,json=debugType,proto3,enum=containerd.runhcs.v1.Options_DebugType" json:"debug_type,omitempty"`
	// registry key root for storage of the runhcs container state
	RegistryRoot string `protobuf:"bytes,3,opt,name=registry_root,json=registryRoot,proto3" json:"registry_root,omitempty"`
	// sandbox_image is the image to use for the sandbox that matches the
	// sandbox_platform.
	SandboxImage string `protobuf:"bytes,4,opt,name=sandbox_image,json=sandboxImage,proto3" json:"sandbox_image,omitempty"`
	// sandbox_platform is a CRI setting that specifies the platform
	// architecture for all sandbox's in this runtime. Values are
	// 'windows/amd64' and 'linux/amd64'.
	SandboxPlatform string `protobuf:"bytes,5,opt,name=sandbox_platform,json=sandboxPlatform,proto3" json:"sandbox_platform,omitempty"`
	// sandbox_isolation is a CRI setting that specifies the isolation level of
	// the sandbox. For Windows runtime PROCESS and HYPERVISOR are valid. For
	// LCOW only HYPERVISOR is valid and default if omitted.
	SandboxIsolation Options_SandboxIsolation `protobuf:"varint,6,opt,name=sandbox_isolation,json=sandboxIsolation,proto3,enum=containerd.runhcs.v1.Options_SandboxIsolation" json:"sandbox_isolation,omitempty"`
	// boot_files_root_path is the path to the directory containing the LCOW
	// kernel and root FS files.
	BootFilesRootPath string `protobuf:"bytes,7,opt,name=boot_files_root_path,json=bootFilesRootPath,proto3" json:"boot_files_root_path,omitempty"`
	// vm_processor_count is the default number of processors to create for the
	// hypervisor isolated utility vm.
	//
	// The platform default if omitted is 2, unless the host only has a single
	// core in which case it is 1.
	VmProcessorCount int32 `protobuf:"varint,8,opt,name=vm_processor_count,json=vmProcessorCount,proto3" json:"vm_processor_count,omitempty"`
	// vm_memory_size_in_mb is the default amount of memory to assign to the
	// hypervisor isolated utility vm.
	//
	// The platform default is 1024MB if omitted.
	VmMemorySizeInMb int32 `protobuf:"varint,9,opt,name=vm_memory_size_in_mb,json=vmMemorySizeInMb,proto3" json:"vm_memory_size_in_mb,omitempty"`
	// GPUVHDPath is the path to the gpu vhd to add to the uvm
	// when a container requests a gpu
	GPUVHDPath string `protobuf:"bytes,10,opt,name=GPUVHDPath,proto3" json:"GPUVHDPath,omitempty"`
	// scale_cpu_limits_to_sandbox indicates that container CPU limits should
	// be adjusted to account for the difference in number of cores between the
	// host and UVM.
	ScaleCpuLimitsToSandbox bool `protobuf:"varint,11,opt,name=scale_cpu_limits_to_sandbox,json=scaleCpuLimitsToSandbox,proto3" json:"scale_cpu_limits_to_sandbox,omitempty"`
	// only_signal_init_process only signals the containers init process during a container
	// kill. If this is off, all processes will be signalled by runc. This is LCOW specific.
	OnlySignalInitProcess bool     `protobuf:"varint,12,opt,name=only_signal_init_process,json=onlySignalInitProcess,proto3" json:"only_signal_init_process,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Options) Reset()      { *m = Options{} }
func (*Options) ProtoMessage() {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_b643df6839c75082, []int{0}
}
func (m *Options) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Options.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return m.Size()
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

// ProcessDetails contains additional information about a process. This is the additional
// info returned in the Pids query.
type ProcessDetails struct {
	ImageName                    string    `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	CreatedAt                    time.Time `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	KernelTime_100Ns             uint64    `protobuf:"varint,3,opt,name=kernel_time_100_ns,json=kernelTime100Ns,proto3" json:"kernel_time_100_ns,omitempty"`
	MemoryCommitBytes            uint64    `protobuf:"varint,4,opt,name=memory_commit_bytes,json=memoryCommitBytes,proto3" json:"memory_commit_bytes,omitempty"`
	MemoryWorkingSetPrivateBytes uint64    `protobuf:"varint,5,opt,name=memory_working_set_private_bytes,json=memoryWorkingSetPrivateBytes,proto3" json:"memory_working_set_private_bytes,omitempty"`
	MemoryWorkingSetSharedBytes  uint64    `protobuf:"varint,6,opt,name=memory_working_set_shared_bytes,json=memoryWorkingSetSharedBytes,proto3" json:"memory_working_set_shared_bytes,omitempty"`
	ProcessID                    uint32    `protobuf:"varint,7,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	UserTime_100Ns               uint64    `protobuf:"varint,8,opt,name=user_time_100_ns,json=userTime100Ns,proto3" json:"user_time_100_ns,omitempty"`
	ExecID                       string    `protobuf:"bytes,9,opt,name=exec_id,json=execId,proto3" json:"exec_id,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}  `json:"-"`
	XXX_unrecognized             []byte    `json:"-"`
	XXX_sizecache                int32     `json:"-"`
}

func (m *ProcessDetails) Reset()      { *m = ProcessDetails{} }
func (*ProcessDetails) ProtoMessage() {}
func (*ProcessDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_b643df6839c75082, []int{1}
}
func (m *ProcessDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessDetails.Merge(m, src)
}
func (m *ProcessDetails) XXX_Size() int {
	return m.Size()
}
func (m *ProcessDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessDetails proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("containerd.runhcs.v1.Options_DebugType", Options_DebugType_name, Options_DebugType_value)
	proto.RegisterEnum("containerd.runhcs.v1.Options_SandboxIsolation", Options_SandboxIsolation_name, Options_SandboxIsolation_value)
	proto.RegisterType((*Options)(nil), "containerd.runhcs.v1.Options")
	proto.RegisterType((*ProcessDetails)(nil), "containerd.runhcs.v1.ProcessDetails")
}

func init() {
	proto.RegisterFile("github.com/Microsoft/hcsshim/cmd/containerd-shim-runhcs-v1/options/runhcs.proto", fileDescriptor_b643df6839c75082)
}

var fileDescriptor_b643df6839c75082 = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x15, 0x63, 0xeb, 0xc1, 0xeb, 0xd8, 0x91, 0xa7, 0x2a, 0x4a, 0x38, 0xad, 0x24, 0x38, 0x8b,
	0x38, 0x68, 0x4c, 0xda, 0xe9, 0xa2, 0x8b, 0x76, 0x53, 0x3d, 0xdc, 0xb0, 0x88, 0x6d, 0x81, 0x72,
	0x93, 0x3e, 0x16, 0x03, 0x8a, 0x1c, 0x53, 0x83, 0x70, 0x38, 0x04, 0x67, 0xa4, 0x5a, 0x59, 0xf5,
	0x13, 0xfa, 0x3f, 0xfd, 0x01, 0x2f, 0xbb, 0x2c, 0x50, 0xc0, 0x6d, 0xf4, 0x25, 0xc5, 0x0c, 0x87,
	0x4e, 0x1b, 0x18, 0xdd, 0x74, 0xa5, 0xe1, 0x39, 0xe7, 0x9e, 0xfb, 0x98, 0x3b, 0x82, 0xf3, 0x84,
	0xca, 0xf9, 0x62, 0xe6, 0x46, 0x9c, 0x79, 0xa7, 0x34, 0x2a, 0xb8, 0xe0, 0x97, 0xd2, 0x9b, 0x47,
	0x42, 0xcc, 0x29, 0xf3, 0x22, 0x16, 0x7b, 0x11, 0xcf, 0x64, 0x48, 0x33, 0x52, 0xc4, 0x87, 0x0a,
	0x3b, 0x2c, 0x16, 0xd9, 0x3c, 0x12, 0x87, 0xcb, 0x63, 0x8f, 0xe7, 0x92, 0xf2, 0x4c, 0x78, 0x25,
	0xe2, 0xe6, 0x05, 0x97, 0x1c, 0x75, 0xde, 0xe9, 0x5d, 0x43, 0x2c, 0x8f, 0xf7, 0x3a, 0x09, 0x4f,
	0xb8, 0x16, 0x78, 0xea, 0x54, 0x6a, 0xf7, 0x7a, 0x09, 0xe7, 0x49, 0x4a, 0x3c, 0xfd, 0x35, 0x5b,
	0x5c, 0x7a, 0x92, 0x32, 0x22, 0x64, 0xc8, 0xf2, 0x52, 0xb0, 0xff, 0x6b, 0x1d, 0x9a, 0xe7, 0x65,
	0x16, 0xd4, 0x81, 0x7a, 0x4c, 0x66, 0x8b, 0xc4, 0xb1, 0xfa, 0xd6, 0x41, 0x2b, 0x28, 0x3f, 0xd0,
	0x09, 0x80, 0x3e, 0x60, 0xb9, 0xca, 0x89, 0x73, 0xaf, 0x6f, 0x1d, 0xec, 0x3c, 0x7b, 0xec, 0xde,
	0x55, 0x83, 0x6b, 0x8c, 0xdc, 0x91, 0xd2, 0x5f, 0xac, 0x72, 0x12, 0xd8, 0x71, 0x75, 0x44, 0x8f,
	0x60, 0xbb, 0x20, 0x09, 0x15, 0xb2, 0x58, 0xe1, 0x82, 0x73, 0xe9, 0x6c, 0xf4, 0xad, 0x03, 0x3b,
	0xb8, 0x5f, 0x81, 0x01, 0xe7, 0x52, 0x89, 0x44, 0x98, 0xc5, 0x33, 0x7e, 0x85, 0x29, 0x0b, 0x13,
	0xe2, 0x6c, 0x96, 0x22, 0x03, 0xfa, 0x0a, 0x43, 0x4f, 0xa0, 0x5d, 0x89, 0xf2, 0x34, 0x94, 0x97,
	0xbc, 0x60, 0x4e, 0x5d, 0xeb, 0x1e, 0x18, 0x7c, 0x62, 0x60, 0xf4, 0x23, 0xec, 0xde, 0xfa, 0x09,
	0x9e, 0x86, 0xaa, 0x3e, 0xa7, 0xa1, 0x7b, 0x70, 0xff, 0xbb, 0x87, 0xa9, 0xc9, 0x58, 0x45, 0x05,
	0x55, 0xce, 0x5b, 0x04, 0x79, 0xd0, 0x99, 0x71, 0x2e, 0xf1, 0x25, 0x4d, 0x89, 0xd0, 0x3d, 0xe1,
	0x3c, 0x94, 0x73, 0xa7, 0xa9, 0x6b, 0xd9, 0x55, 0xdc, 0x89, 0xa2, 0x54, 0x67, 0x93, 0x50, 0xce,
	0xd1, 0x53, 0x40, 0x4b, 0x86, 0xf3, 0x82, 0x47, 0x44, 0x08, 0x5e, 0xe0, 0x88, 0x2f, 0x32, 0xe9,
	0xb4, 0xfa, 0xd6, 0x41, 0x3d, 0x68, 0x2f, 0xd9, 0xa4, 0x22, 0x86, 0x0a, 0x47, 0x2e, 0x74, 0x96,
	0x0c, 0x33, 0xc2, 0x78, 0xb1, 0xc2, 0x82, 0xbe, 0x21, 0x98, 0x66, 0x98, 0xcd, 0x1c, 0xbb, 0xd2,
	0x9f, 0x6a, 0x6a, 0x4a, 0xdf, 0x10, 0x3f, 0x3b, 0x9d, 0xa1, 0x2e, 0xc0, 0xd7, 0x93, 0x6f, 0x5f,
	0x3e, 0x1f, 0xa9, 0x5c, 0x0e, 0xe8, 0x22, 0xfe, 0x81, 0xa0, 0x2f, 0xe1, 0xa1, 0x88, 0xc2, 0x94,
	0xe0, 0x28, 0x5f, 0xe0, 0x94, 0x32, 0x2a, 0x05, 0x96, 0x1c, 0x9b, 0xb6, 0x9c, 0x2d, 0x7d, 0xe9,
	0x1f, 0x69, 0xc9, 0x30, 0x5f, 0xbc, 0xd0, 0x82, 0x0b, 0x6e, 0xe6, 0x80, 0x3e, 0x07, 0x87, 0x67,
	0xa9, 0x2a, 0x24, 0xc9, 0xc2, 0x14, 0xd3, 0x8c, 0xca, 0xaa, 0x13, 0xe7, 0xbe, 0x0e, 0xfd, 0x50,
	0xf1, 0x53, 0x4d, 0xfb, 0x19, 0x95, 0xa6, 0x9b, 0xfd, 0x27, 0x60, 0xdf, 0xee, 0x03, 0xb2, 0xa1,
	0x7e, 0x36, 0xf1, 0x27, 0xe3, 0x76, 0x0d, 0xb5, 0x60, 0xf3, 0xc4, 0x7f, 0x31, 0x6e, 0x5b, 0xa8,
	0x09, 0x1b, 0xe3, 0x8b, 0x57, 0xed, 0x7b, 0xfb, 0x1e, 0xb4, 0xdf, 0x1f, 0x3b, 0xda, 0x82, 0xe6,
	0x24, 0x38, 0x1f, 0x8e, 0xa7, 0xd3, 0x76, 0x0d, 0xed, 0x00, 0x3c, 0xff, 0x7e, 0x32, 0x0e, 0x5e,
	0xfa, 0xd3, 0xf3, 0xa0, 0x6d, 0xed, 0xff, 0xb1, 0x01, 0x3b, 0x26, 0xcf, 0x88, 0xc8, 0x90, 0xa6,
	0x02, 0x7d, 0x02, 0xa0, 0x37, 0x07, 0x67, 0x21, 0x23, 0x7a, 0x93, 0xed, 0xc0, 0xd6, 0xc8, 0x59,
	0xc8, 0x08, 0x1a, 0x02, 0x44, 0x05, 0x09, 0x25, 0x89, 0x71, 0x28, 0xf5, 0x36, 0x6f, 0x3d, 0xdb,
	0x73, 0xcb, 0x57, 0xe2, 0x56, 0xaf, 0xc4, 0xbd, 0xa8, 0x5e, 0xc9, 0xa0, 0x75, 0x7d, 0xd3, 0xab,
	0xfd, 0xf2, 0x67, 0xcf, 0x0a, 0x6c, 0x13, 0xf7, 0x95, 0x44, 0x9f, 0x02, 0x7a, 0x4d, 0x8a, 0x8c,
	0xa4, 0x58, 0x3d, 0x27, 0x7c, 0x7c, 0x74, 0x84, 0x33, 0xa1, 0xf7, 0x79, 0x33, 0x78, 0x50, 0x32,
	0xca, 0xe1, 0xf8, 0xe8, 0xe8, 0x4c, 0x20, 0x17, 0x3e, 0x30, 0x77, 0x18, 0x71, 0xc6, 0xa8, 0xc4,
	0xb3, 0x95, 0x24, 0x42, 0x2f, 0xf6, 0x66, 0xb0, 0x5b, 0x52, 0x43, 0xcd, 0x0c, 0x14, 0x81, 0x4e,
	0xa0, 0x6f, 0xf4, 0x3f, 0xf1, 0xe2, 0x35, 0xcd, 0x12, 0x2c, 0x88, 0x1a, 0x35, 0x5d, 0x86, 0x92,
	0x98, 0xe0, 0xba, 0x0e, 0xfe, 0xb8, 0xd4, 0xbd, 0x2a, 0x65, 0x53, 0x22, 0x27, 0xa5, 0xa8, 0xf4,
	0x19, 0x41, 0xef, 0x0e, 0x1f, 0x31, 0x0f, 0x0b, 0x12, 0x1b, 0x9b, 0x86, 0xb6, 0x79, 0xf8, 0xbe,
	0xcd, 0x54, 0x6b, 0x4a, 0x97, 0xa7, 0x00, 0xe6, 0x96, 0x31, 0x8d, 0xf5, 0x66, 0x6f, 0x0f, 0xb6,
	0xd7, 0x37, 0x3d, 0xdb, 0x8c, 0xdd, 0x1f, 0x05, 0xb6, 0x11, 0xf8, 0x31, 0x7a, 0x0c, 0xed, 0x85,
	0x20, 0xc5, 0xbf, 0xc6, 0xd2, 0xd2, 0x49, 0xb6, 0x15, 0xfe, 0x6e, 0x28, 0x8f, 0xa0, 0x49, 0xae,
	0x48, 0xa4, 0x3c, 0xd5, 0x3a, 0xdb, 0x03, 0x58, 0xdf, 0xf4, 0x1a, 0xe3, 0x2b, 0x12, 0xf9, 0xa3,
	0xa0, 0xa1, 0x28, 0x3f, 0x1e, 0xc4, 0xd7, 0x6f, 0xbb, 0xb5, 0xdf, 0xdf, 0x76, 0x6b, 0x3f, 0xaf,
	0xbb, 0xd6, 0xf5, 0xba, 0x6b, 0xfd, 0xb6, 0xee, 0x5a, 0x7f, 0xad, 0xbb, 0xd6, 0x0f, 0xdf, 0xfc,
	0xff, 0xff, 0xd4, 0x2f, 0xcc, 0xef, 0x77, 0xb5, 0x59, 0x43, 0xdf, 0xfb, 0x67, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0xdb, 0x30, 0x52, 0xaa, 0x05, 0x00, 0x00,
}

func (m *Options) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Options) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Debug {
		dAtA[i] = 0x8
		i++
		if m.Debug {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DebugType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(m.DebugType))
	}
	if len(m.RegistryRoot) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(len(m.RegistryRoot)))
		i += copy(dAtA[i:], m.RegistryRoot)
	}
	if len(m.SandboxImage) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(len(m.SandboxImage)))
		i += copy(dAtA[i:], m.SandboxImage)
	}
	if len(m.SandboxPlatform) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(len(m.SandboxPlatform)))
		i += copy(dAtA[i:], m.SandboxPlatform)
	}
	if m.SandboxIsolation != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(m.SandboxIsolation))
	}
	if len(m.BootFilesRootPath) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(len(m.BootFilesRootPath)))
		i += copy(dAtA[i:], m.BootFilesRootPath)
	}
	if m.VmProcessorCount != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(m.VmProcessorCount))
	}
	if m.VmMemorySizeInMb != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(m.VmMemorySizeInMb))
	}
	if len(m.GPUVHDPath) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(len(m.GPUVHDPath)))
		i += copy(dAtA[i:], m.GPUVHDPath)
	}
	if m.ScaleCpuLimitsToSandbox {
		dAtA[i] = 0x58
		i++
		if m.ScaleCpuLimitsToSandbox {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.OnlySignalInitProcess {
		dAtA[i] = 0x60
		i++
		if m.OnlySignalInitProcess {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ProcessDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ImageName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(len(m.ImageName)))
		i += copy(dAtA[i:], m.ImageName)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintRunhcs(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.KernelTime_100Ns != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(m.KernelTime_100Ns))
	}
	if m.MemoryCommitBytes != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(m.MemoryCommitBytes))
	}
	if m.MemoryWorkingSetPrivateBytes != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(m.MemoryWorkingSetPrivateBytes))
	}
	if m.MemoryWorkingSetSharedBytes != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(m.MemoryWorkingSetSharedBytes))
	}
	if m.ProcessID != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(m.ProcessID))
	}
	if m.UserTime_100Ns != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(m.UserTime_100Ns))
	}
	if len(m.ExecID) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintRunhcs(dAtA, i, uint64(len(m.ExecID)))
		i += copy(dAtA[i:], m.ExecID)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRunhcs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Options) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Debug {
		n += 2
	}
	if m.DebugType != 0 {
		n += 1 + sovRunhcs(uint64(m.DebugType))
	}
	l = len(m.RegistryRoot)
	if l > 0 {
		n += 1 + l + sovRunhcs(uint64(l))
	}
	l = len(m.SandboxImage)
	if l > 0 {
		n += 1 + l + sovRunhcs(uint64(l))
	}
	l = len(m.SandboxPlatform)
	if l > 0 {
		n += 1 + l + sovRunhcs(uint64(l))
	}
	if m.SandboxIsolation != 0 {
		n += 1 + sovRunhcs(uint64(m.SandboxIsolation))
	}
	l = len(m.BootFilesRootPath)
	if l > 0 {
		n += 1 + l + sovRunhcs(uint64(l))
	}
	if m.VmProcessorCount != 0 {
		n += 1 + sovRunhcs(uint64(m.VmProcessorCount))
	}
	if m.VmMemorySizeInMb != 0 {
		n += 1 + sovRunhcs(uint64(m.VmMemorySizeInMb))
	}
	l = len(m.GPUVHDPath)
	if l > 0 {
		n += 1 + l + sovRunhcs(uint64(l))
	}
	if m.ScaleCpuLimitsToSandbox {
		n += 2
	}
	if m.OnlySignalInitProcess {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProcessDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ImageName)
	if l > 0 {
		n += 1 + l + sovRunhcs(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovRunhcs(uint64(l))
	if m.KernelTime_100Ns != 0 {
		n += 1 + sovRunhcs(uint64(m.KernelTime_100Ns))
	}
	if m.MemoryCommitBytes != 0 {
		n += 1 + sovRunhcs(uint64(m.MemoryCommitBytes))
	}
	if m.MemoryWorkingSetPrivateBytes != 0 {
		n += 1 + sovRunhcs(uint64(m.MemoryWorkingSetPrivateBytes))
	}
	if m.MemoryWorkingSetSharedBytes != 0 {
		n += 1 + sovRunhcs(uint64(m.MemoryWorkingSetSharedBytes))
	}
	if m.ProcessID != 0 {
		n += 1 + sovRunhcs(uint64(m.ProcessID))
	}
	if m.UserTime_100Ns != 0 {
		n += 1 + sovRunhcs(uint64(m.UserTime_100Ns))
	}
	l = len(m.ExecID)
	if l > 0 {
		n += 1 + l + sovRunhcs(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRunhcs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRunhcs(x uint64) (n int) {
	return sovRunhcs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Options) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Options{`,
		`Debug:` + fmt.Sprintf("%v", this.Debug) + `,`,
		`DebugType:` + fmt.Sprintf("%v", this.DebugType) + `,`,
		`RegistryRoot:` + fmt.Sprintf("%v", this.RegistryRoot) + `,`,
		`SandboxImage:` + fmt.Sprintf("%v", this.SandboxImage) + `,`,
		`SandboxPlatform:` + fmt.Sprintf("%v", this.SandboxPlatform) + `,`,
		`SandboxIsolation:` + fmt.Sprintf("%v", this.SandboxIsolation) + `,`,
		`BootFilesRootPath:` + fmt.Sprintf("%v", this.BootFilesRootPath) + `,`,
		`VmProcessorCount:` + fmt.Sprintf("%v", this.VmProcessorCount) + `,`,
		`VmMemorySizeInMb:` + fmt.Sprintf("%v", this.VmMemorySizeInMb) + `,`,
		`GPUVHDPath:` + fmt.Sprintf("%v", this.GPUVHDPath) + `,`,
		`ScaleCpuLimitsToSandbox:` + fmt.Sprintf("%v", this.ScaleCpuLimitsToSandbox) + `,`,
		`OnlySignalInitProcess:` + fmt.Sprintf("%v", this.OnlySignalInitProcess) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessDetails) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessDetails{`,
		`ImageName:` + fmt.Sprintf("%v", this.ImageName) + `,`,
		`CreatedAt:` + strings.Replace(strings.Replace(this.CreatedAt.String(), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`KernelTime_100Ns:` + fmt.Sprintf("%v", this.KernelTime_100Ns) + `,`,
		`MemoryCommitBytes:` + fmt.Sprintf("%v", this.MemoryCommitBytes) + `,`,
		`MemoryWorkingSetPrivateBytes:` + fmt.Sprintf("%v", this.MemoryWorkingSetPrivateBytes) + `,`,
		`MemoryWorkingSetSharedBytes:` + fmt.Sprintf("%v", this.MemoryWorkingSetSharedBytes) + `,`,
		`ProcessID:` + fmt.Sprintf("%v", this.ProcessID) + `,`,
		`UserTime_100Ns:` + fmt.Sprintf("%v", this.UserTime_100Ns) + `,`,
		`ExecID:` + fmt.Sprintf("%v", this.ExecID) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRunhcs(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Options) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRunhcs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Options: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Options: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debug", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Debug = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebugType", wireType)
			}
			m.DebugType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DebugType |= Options_DebugType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryRoot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunhcs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunhcs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegistryRoot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SandboxImage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunhcs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunhcs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SandboxImage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SandboxPlatform", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunhcs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunhcs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SandboxPlatform = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SandboxIsolation", wireType)
			}
			m.SandboxIsolation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SandboxIsolation |= Options_SandboxIsolation(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BootFilesRootPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunhcs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunhcs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BootFilesRootPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VmProcessorCount", wireType)
			}
			m.VmProcessorCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VmProcessorCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VmMemorySizeInMb", wireType)
			}
			m.VmMemorySizeInMb = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VmMemorySizeInMb |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GPUVHDPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunhcs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunhcs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GPUVHDPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScaleCpuLimitsToSandbox", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ScaleCpuLimitsToSandbox = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnlySignalInitProcess", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OnlySignalInitProcess = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRunhcs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRunhcs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRunhcs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProcessDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRunhcs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProcessDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunhcs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunhcs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRunhcs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRunhcs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KernelTime_100Ns", wireType)
			}
			m.KernelTime_100Ns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KernelTime_100Ns |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryCommitBytes", wireType)
			}
			m.MemoryCommitBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryCommitBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryWorkingSetPrivateBytes", wireType)
			}
			m.MemoryWorkingSetPrivateBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryWorkingSetPrivateBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryWorkingSetSharedBytes", wireType)
			}
			m.MemoryWorkingSetSharedBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryWorkingSetSharedBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessID", wireType)
			}
			m.ProcessID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProcessID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserTime_100Ns", wireType)
			}
			m.UserTime_100Ns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserTime_100Ns |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunhcs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunhcs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRunhcs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRunhcs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRunhcs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRunhcs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRunhcs
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRunhcs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRunhcs
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRunhcs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRunhcs
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRunhcs(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRunhcs
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRunhcs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRunhcs   = fmt.Errorf("proto: integer overflow")
)
