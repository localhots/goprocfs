package procfs

import (
	"fmt"
	"io/ioutil"
)

// ProcPidStatm contains process information obtained from /proc/[PID]/statm
// Its fields are documented using PROC(5) man page
type ProcPidStatm struct {
	Size     uint64 // Total program size
	Resident uint64 // Resident set size
	Share    uint64 // Shared pages
	Text     uint64 // Text (code)
	Lib      uint64 // Library (unused in Linux 2.6)
	Data     uint64 // Data + Stack
	Dt       uint64 // Dirty pages (unused in Linux 2.6)
}

func NewProcPidStatm(pid int) (ProcPidStatm, error) {
	p := ProcPidStatm{}

	file := fmt.Sprintf("/proc/%d/statm", pid)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return p, err
	}

	parsed, err := fmt.Sscanf(string(b), "%d %d %d %d %d %d %d",
		&p.Size, &p.Resident, &p.Share, &p.Text, &p.Lib, &p.Data, &p.Dt)

	if parsed < 7 {
		err := fmt.Errorf("Managed to parse only %d fields out of 7", parsed)
		return p, err
	}
	if err != nil {
		return p, err
	}

	return p, nil
}
