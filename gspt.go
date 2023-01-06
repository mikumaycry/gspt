package gspt

const (
	// These values must match the return values for spt_init1() used in C.
	HaveNone        = 0
	HaveNative      = 1
	HaveReplacement = 2
)

var (
	HaveSetProcTitle     int
	HaveSetProcTitleFast int
)
