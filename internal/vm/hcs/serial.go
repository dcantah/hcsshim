package hcs

import (
	"strconv"
	"strings"

	hcsschema "github.com/Microsoft/hcsshim/internal/schema2"
	"github.com/Microsoft/hcsshim/internal/vm"
	"github.com/pkg/errors"
)

func (uvm *utilityVM) SetSerialConsole(port uint32, listenerPath string) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}

	if !strings.HasPrefix(listenerPath, `\\.\pipe\`) {
		return errors.New("listener for serial console is not a named pipe")
	}

	uvm.doc.VirtualMachine.Devices.ComPorts = map[string]hcsschema.ComPort{
		strconv.Itoa(int(port)): { // "0" would be COM1
			NamedPipe: listenerPath,
		},
	}
	return nil
}
