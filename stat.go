package procfs

import (
	"fmt"
	"io/ioutil"
)

// Proc contains process information obtained from /proc/stat
// Its fields are documented using PROC(5) man page
type ProcStat struct {
	Cpu struct {
		User    uint64 // Time spent in user mode
		Nice    uint64 // Time spent in user mode with low priority (nice)
		System  uint64 // Time spent in system mode
		Idle    uint64 // Time spent in the idle task
		Iowait  uint64 // Time waiting for I/O to complete
		Irq     uint64 // Time servicing interrupts
		Softirq uint64 // Time servicing softirqs

		// Stolen time, which is the time spent in other operating systems when
		// running in a virtualized environment
		Steal uint64
		// Time spent running a virtual CPU for guest operating systems under
		// the control of the Linux kernel.
		Guest uint64
		// Time spent running a niced guest (virtual CPU for guest operating
		// systems under the control of the Linux kernel)
		GuestNice uint64
	}
}

func NewProcStat() (ProcStat, error) {
	p := ProcStat{}

	b, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return p, err
	}

	parsed, err := fmt.Sscanf(string(b), "cpu  %d %d %d %d %d %d %d %d %d %d",
		&p.Cpu.User, &p.Cpu.Nice, &p.Cpu.System, &p.Cpu.Idle, &p.Cpu.Iowait,
		&p.Cpu.Irq, &p.Cpu.Softirq, &p.Cpu.Steal, &p.Cpu.Guest,
		&p.Cpu.GuestNice)
	if parsed < 10 {
		err := fmt.Errorf("Managed to parse only %d fields out of 10", parsed)
		return p, err
	}
	if err != nil {
		return p, err
	}

	return p, nil
}
