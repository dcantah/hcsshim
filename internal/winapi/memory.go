package winapi

// VOID RtlMoveMemory(
// 	_Out_       VOID UNALIGNED *Destination,
// 	_In_  const VOID UNALIGNED *Source,
// 	_In_        SIZE_T         Length
// );
//sys RtlMoveMemory(destination *byte, source *byte, length uintptr) (err error) = kernel32.RtlMoveMemory

//sys LocalAlloc(flags uint32, size int) (ptr uintptr) = kernel32.LocalAlloc
//sys LocalFree(ptr uintptr) = kernel32.LocalFree

// HANDLE GetProcessHeap();
//
//sys GetProcessHeap() (procHeap windows.Handle, err error) = kernel32.GetProcessHeap

// DECLSPEC_ALLOCATOR LPVOID HeapAlloc(
// 	HANDLE hHeap,
// 	DWORD  dwFlags,
// 	SIZE_T dwBytes
// );
//
//sys HeapAlloc(hHeap windows.Handle, dwFlags uint32, dwBytes uintptr) (lpMem uintptr, err error) = kernel32.HeapAlloc

// BOOL HeapFree(
// 	HANDLE                 hHeap,
// 	DWORD                  dwFlags,
// 	_Frees_ptr_opt_ LPVOID lpMem
// );
//
//sys HeapFree(hHeap windows.Handle, dwFlags uint32, lpMem uintptr) (err error) = kernel32.HeapFree
