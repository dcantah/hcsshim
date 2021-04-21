package remotevm

import (
	"github.com/Microsoft/hcsshim/internal/vm"
	"github.com/pkg/errors"
)

func (uvm *remoteVM) SetStorageQos(iopsMaximum int64, bandwidthMaximum int64) error {
	if uvm.state != vm.StatePreCreated {
		return vm.ErrNotInPreCreatedState
	}

	// The way that HCS handles these options is a bit odd. They launch the vmworker process in a job object and
	// set the bandwidth and iops limits on the worker process' job object. To keep parity with what we expose today
	// in HCS we can do the same here as we launch the server process in a job object.
	if uvm.job != nil {
		if err := uvm.job.SetIOLimit(bandwidthMaximum, iopsMaximum); err != nil {
			return errors.Wrap(err, "failed to set storage qos values on remotevm process")
		}
	}

	return nil
}
