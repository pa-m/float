// +build !amd64

package float

import (
	"testing"

	"golang.org/x/sys/cpu"
)

func TestHasAVX(t *testing.T) {
	if hasAVX() != cpu.X86.HasAVX2 {
		t.Fail()
	}
}
