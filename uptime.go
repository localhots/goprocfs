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

func NewProcUptime() (ProcUptime, error) {
	p := ProcUptime{}

	b, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return p, err
	}

	parsed, err := fmt.Sscanf(string(b), "%f %f", &p.Uptime, &p.Idle)
	if parsed < 2 {
		err := fmt.Errorf("Managed to parse only %d fields out of 2", parsed)
		return p, err
	}
	if err != nil {
		return p, err
	}

	return p, nil
}
