package jobcontainers

import (
	"fmt"
	"strings"

	"github.com/Microsoft/hcsshim/internal/winapi"
	"golang.org/x/sys/windows"
)

// Takes in a DOMAIN\Username or just Username and will return a token
// for the account if successful.
func processToken(user string) (windows.Token, error) {
	var (
		domain   string
		userName string
		token    windows.Token
	)

	split := strings.Split(user, "\\")
	if len(split) == 2 {
		domain = split[0]
		userName = split[1]
	} else if len(split) == 1 {
		userName = split[0]
	}

	// If empty or ContainerUser or ContainerAdministrator just let it inherit the token
	// from whatever is used to launch it (containerd-shim etc). Regular container images
	// are usable for jobcontainers so we need to handle the ContainerUser and ContainerAdministrator
	// cases also.
	if user == "" || user == "ContainerUser" || user == "ContainerAdministrator" {
		return openCurrentProcessToken()
	}

	var logonType uint32
	// User asking to run as a local system account (NETWORK SERVICE, LOCAL SERVICE, SYSTEM)
	if domain == "NT AUTHORITY" {
		logonType = winapi.LOGON32_LOGON_SERVICE
	} else {
		// They want a user account, use the interactive logon type instead of service
		logonType = winapi.LOGON32_LOGON_INTERACTIVE
	}

	if err := winapi.LogonUser(
		windows.StringToUTF16Ptr(userName),
		windows.StringToUTF16Ptr(domain),
		nil,
		logonType,
		winapi.LOGON32_PROVIDER_DEFAULT,
		&token,
	); err != nil {
		return token, fmt.Errorf("failed to logon user: %s", err)
	}
	return token, nil
}

func openCurrentProcessToken() (windows.Token, error) {
	var token windows.Token
	if err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_ALL_ACCESS, &token); err != nil {
		return 0, err
	}
	return token, nil
}
