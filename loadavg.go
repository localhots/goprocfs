package procfs

import (
	"fmt"
	"io/ioutil"
)

// ProcLoadavg contains process information obtained from /proc/loadavg
// Its fields are documented using PROC(5) man page
type ProcLoadavg struct {
	// The first three fields in this file are load average figures giving the
	// number of jobs in the run queue (state R) or waiting for disk I/O
	// (state D) averaged over 1, 5, and 15 minutes

	Avg1Min  float32 // Average over 1 minute
	Avg5Min  float32 // Average over 5 minutes
	Avg15Min float32 // Average over 15 minutes

	// Number of currently runnable kernel scheduling entities (processes,
	// threads)
	RunnableEntities uint
	// Number of kernel scheduling entities that currently exist on the system
	TotalEntities uint
	// PID of the process that was most recently created on the system
	LastPid uint
}

func NewProcLoadavg() ProcLoadavg {
	b, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		panic(err)
	}

	p := ProcLoadavg{}

	parsed, err := fmt.Sscanf(string(b), "%f %f %f %d/%d %d",
		&p.Avg1Min, &p.Avg5Min, &p.Avg15Min, &p.RunnableEntities,
		&p.TotalEntities, &p.LastPid)
	if parsed < 6 {
		fmt.Println("Managed to parse only", parsed, "fields out of 6")
	}
	if err != nil {
		panic(err)
	}

	return p
}
