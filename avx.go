package float

import "golang.org/x/sys/cpu"

func hasAVX() bool { return cpu.X86.HasAVX2 }
