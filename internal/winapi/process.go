package winapi

const PROCESS_ALL_ACCESS uint32 = 0x1FFFFF

const (
	PROC_THREAD_ATTRIBUTE_JOB_LIST       uintptr = 0x2000D
	PROC_THREAD_ATTRIBUTE_PSEUDOCONSOLE  uintptr = 0x20016
	PROC_THREAD_ATTRIBUTE_PARENT_PROCESS uintptr = 0x20000
)

// BOOL CreateProcessAsUserW(
// 	HANDLE                hToken,
// 	LPCWSTR               lpApplicationName,
// 	LPWSTR                lpCommandLine,
// 	LPSECURITY_ATTRIBUTES lpProcessAttributes,
// 	LPSECURITY_ATTRIBUTES lpThreadAttributes,
// 	BOOL                  bInheritHandles,
// 	DWORD                 dwCreationFlags,
// 	LPVOID                lpEnvironment,
// 	LPCWSTR               lpCurrentDirectory,
// 	LPSTARTUPINFOW        lpStartupInfo,
// 	LPPROCESS_INFORMATION lpProcessInformation
// );
//sys CreateProcessAsUser(hToken windows.Token, appName *uint16, commandLine *uint16, procSecurity *windows.SecurityAttributes, threadSecurity *windows.SecurityAttributes, inheritHandles bool, creationFlags uint32, env *uint16, currentDir *uint16, startupInfo *windows.StartupInfo, outProcInfo *windows.ProcessInformation) (err error) = Advapi32.CreateProcessAsUserW

// BOOL InitializeProcThreadAttributeList(
// 	LPPROC_THREAD_ATTRIBUTE_LIST lpAttributeList,
// 	DWORD                        dwAttributeCount,
// 	DWORD                        dwFlags,
// 	PSIZE_T                      lpSize
// );
//
//sys InitializeProcThreadAttributeList(lpAttributeList uintptr, dwAttributeCount uint32, dwFlags uint32, lpSize *uintptr) (err error) = kernel32.InitializeProcThreadAttributeList

// BOOL UpdateProcThreadAttribute(
// 	LPPROC_THREAD_ATTRIBUTE_LIST lpAttributeList,
// 	DWORD                        dwFlags,
// 	DWORD_PTR                    Attribute,
// 	PVOID                        lpValue,
// 	SIZE_T                       cbSize,
// 	PVOID                        lpPreviousValue,
// 	PSIZE_T                      lpReturnSize
// );
//
//sys UpdateProcThreadAttribute(lpAttributeList uintptr, dwFlags uint32, attribute uintptr, lpValue *uintptr, cbSize uintptr, lpPreviousValue *uintptr, lpReturnSize *uintptr) (err error) = kernel32.UpdateProcThreadAttribute

// void DeleteProcThreadAttributeList(
// 	LPPROC_THREAD_ATTRIBUTE_LIST lpAttributeList
// );
//
//sys DeleteProcThreadAttributeList(lpAttributeList uintptr) = kernel32.DeleteProcThreadAttributeList
