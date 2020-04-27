package processorinfo

import (
	"runtime"
	"syscall"
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-getmaximumprocessorcount
	getMaxProcCount = modkernel32.NewProc("GetMaximumProcessorCount")
)

// Get count from all processor groups.
// https://docs.microsoft.com/en-us/windows/win32/procthread/processor-groups
const ALL_PROCESSOR_GROUPS = 0xFFFF

// ProcessorCount calls the win32 API function GetMaximumProcessorCount
// to get the total number of logical processors on the system. If this
// fails it will fall back to runtime.NumCPU
func ProcessorCount() int {
	ret, _, _ := getMaxProcCount.Call(ALL_PROCESSOR_GROUPS)
	// If the call fails the return value will be 0. If so fall
	// back to NumCPU
	if ret == 0 {
		return runtime.NumCPU()
	}
	return int(ret)
}
