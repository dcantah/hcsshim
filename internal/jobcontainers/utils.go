package jobcontainers

import (
	"strings"

	"github.com/Microsoft/hcsshim/internal/jobobject"
)

// Seperates path to executable/cmd from it's arguments. The path itself needs to be the
// first element in the arguments.
func separateArgs(cmdline string) (string, []string) {
	split := strings.Fields(cmdline)
	return split[0], split
}

func calculateJobCPUWeight(processorWeight uint32) uint32 {
	return 1 + uint32((8*processorWeight)/jobobject.CPUWeightMax)
}

func calculateJobCPURate(hostProcs uint32, processorCount uint32) uint32 {
	rate := (processorCount * 10000) / hostProcs
	if rate == 0 {
		return 1
	}
	return rate
}

// Convert environment map to a slice of environment variables in the form [Key1=val1, key2=val2]
func envMapToSlice(m map[string]string) []string {
	var s []string
	for k, v := range m {
		s = append(s, k+"="+v)
	}
	return s
}
