package privileged

import (
	"strconv"
	"testing"
)

func TestBitmask(t *testing.T) {
	mask := int32ToBitmask(5)
	str := strconv.FormatInt(int64(mask), 2)
	if str != "11111" {
		t.Fatalf("expected '11111' but got: %s", str)
	}
}
