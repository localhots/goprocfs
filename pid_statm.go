package goprocfs

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

func NewProcPidStatm(pid int) ProcPidStatm {
	file := fmt.Sprintf("/proc/%d/statm", pid)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	p := ProcPidStatm{}

	parsed, err := fmt.Sscanf(string(bytes), "%d %d %d %d %d %d %d",
		&p.Size, &p.Resident, &p.Share, &p.Text, &p.Lib, &p.Data, &p.Dt)

	if parsed < 7 {
		fmt.Println("Managed to parse only", parsed, "fields out of 7")
	}
	if err != nil {
		panic(err)
	}

	return p
}
