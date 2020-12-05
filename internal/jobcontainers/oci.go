package jobcontainers

import (
	"context"
	"fmt"

	"github.com/Microsoft/hcsshim/internal/jobobject"
	"github.com/Microsoft/hcsshim/internal/processorinfo"

	"github.com/Microsoft/hcsshim/internal/log"
	"github.com/Microsoft/hcsshim/internal/oci"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
)

// This file contains helpers for converting parts of the oci spec to useful
// structures/limits to be applied to a job object.

// Oci spec to job object limit information. Will do any conversions to job object specific values from
// their respective OCI representations. E.g. we convert CPU count into the correct job object cpu
// rate value internally.
func specToLimits(ctx context.Context, s *specs.Spec) (*jobobject.JobLimits, error) {
	// CPU limits
	cpuNumSet := 0
	cpuCount := uint32(oci.ParseAnnotationsCPUCount(ctx, s, oci.AnnotationContainerProcessorCount, 0))
	if cpuCount > 0 {
		cpuNumSet++
	}

	cpuLimit := uint32(oci.ParseAnnotationsCPULimit(ctx, s, oci.AnnotationContainerProcessorLimit, 0))
	if cpuLimit > 0 {
		cpuNumSet++
	}

	cpuWeight := uint32(oci.ParseAnnotationsCPUWeight(ctx, s, oci.AnnotationContainerProcessorWeight, 0))
	if cpuWeight > 0 {
		cpuNumSet++
	}

	if cpuNumSet > 1 {
		return nil, fmt.Errorf("invalid spec - Windows Job Container CPU Count: '%d', Limit: '%d', and Weight: '%d' are mutually exclusive", cpuCount, cpuLimit, cpuWeight)
	} else if cpuNumSet == 1 {
		if cpuCount != 0 {
			hostCPUCount := uint32(processorinfo.ProcessorCount())
			if cpuCount > hostCPUCount {
				log.G(ctx).WithFields(logrus.Fields{
					"requested": cpuCount,
					"assigned":  hostCPUCount,
				}).Warn("Changing user requested CPUCount to current number of processors")
				cpuCount = hostCPUCount
			}
			// Job object API does not support "CPU count". Instead, we translate the notion of "count" into
			// CPU limit, which represents the amount of the host system's processors that the job can use to
			// a percentage times 100. For example, to let the job use 20% of the available LPs the rate would
			// be 20 times 100, or 2,000.
			cpuLimit = calculateJobCPURate(hostCPUCount, cpuCount)
		} else if cpuWeight != 0 {
			cpuWeight = calculateJobCPUWeight(cpuWeight)
		}
		// Nothing to do for cpu limit, we can assign directly.
	}

	// Memory limit
	memLimit := oci.ParseAnnotationsMemory(ctx, s, oci.AnnotationContainerMemorySizeInMB, 0)

	// IO limits
	maxBandwidth := int64(oci.ParseAnnotationsStorageBps(ctx, s, oci.AnnotationContainerStorageQoSBandwidthMaximum, 0))
	maxIops := int64(oci.ParseAnnotationsStorageIops(ctx, s, oci.AnnotationContainerStorageQoSIopsMaximum, 0))

	return &jobobject.JobLimits{
		CPULimit:           cpuLimit,
		CPUWeight:          cpuWeight,
		MaxIOPS:            maxIops,
		MaxBandwidth:       maxBandwidth,
		MemoryLimitInBytes: memLimit * 1024 * 1024, // ParseAnnotationsMemory value is returned in MB
	}, nil
}
