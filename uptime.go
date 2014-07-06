package procfs

import (
	"fmt"
	"io/ioutil"
)

// Proc contains process information obtained from /proc/uptime
// Its fields are documented using PROC(5) man page
type ProcUptime struct {
	Uptime float64 // Uptime of the system in seconds
	Idle   float64 // Time spent in idle process in seconds
}

func NewProcUptime() ProcUptime {
	b, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		panic(err)
	}

	p := ProcUptime{}

	parsed, err := fmt.Sscanf(string(b), "%f %f", &p.Uptime, &p.Idle)
	if parsed < 2 {
		fmt.Println("Managed to parse only", parsed, "fields out of 2")
	}
	if err != nil {
		panic(err)
	}

	return p
}
