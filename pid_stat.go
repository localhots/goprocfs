package goprocfs

import (
	"fmt"
	"io/ioutil"
)

// ProcPidStat contains process information obtained from /proc/[PID]/stat
// Its fields are documented using PROC(5) man page
type ProcPidStat struct {
	// The process ID.
	Pid int

	// The filename of the executable, in parentheses.  This is visible whether
	// or not the executable is swapped out.
	Comm string

	// One character from the string "RSDZTW" where R is running, S is sleeping
	// in an interruptible wait, D is waiting in uninterruptible disk sleep, Z
	// is zombie, T is traced or stopped (on a signal), and W is paging.
	State string

	// The PID of the parent.
	Ppid int

	// The process group ID of the process.
	Pgrp int

	// The session ID of the process.
	Session int

	// The controlling terminal of the process. (The minor device number is
	// contained in the combination of bits 31 to 20 and 7 to 0; the major
	// device number is in bits 15 to 8.)
	TtyNr int

	// The ID of the foreground process group of the controlling terminal of
	// the process.
	Tpgid int

	// The kernel flags word of the process. For bit meanings, see the PF_*
	// defines in the Linux kernel source file include/linux/sched.h. Details
	// depend on the kernel version.
	// uint32 before Linux 2.6.22
	Flags uint

	// The number of minor faults the process has made which have not required
	// loading a memory page from disk.
	Minflt uint32

	// The number of minor faults that the process's waited-for children have
	// made.
	Cminflt uint32

	// The number of major faults the process has made which have required
	// loading a memory page from disk.
	Majflt uint32

	// The number of major faults that the process's waited-for children have
	// made.
	Cmajflt uint32

	// Amount of time that this process has been scheduled in user mode,
	// measured in clock ticks (divide by sysconf(_SC_CLK_TCK)). This includes
	// guest time, guest_time (time spent running a virtual CPU, see below), so
	// that applications that are not aware of the guest time field do not lose
	// that time from their calculations.
	Utime uint32

	// Amount of time that this process has been scheduled in kernel mode,
	// measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	Stime uint32

	// Amount of time that this process's waited-for children have been
	// scheduled in user mode, measured in clock ticks (divide by
	// sysconf(_SC_CLK_TCK)). (See also times(2).) This includes guest time,
	// cguest_time (time spent running a virtual CPU, see below).
	Cutime int32

	// Amount of time that this process's waited-for children have been
	// scheduled in kernel mode, measured in clock ticks (divide by
	// sysconf(_SC_CLK_TCK)).
	Cstime int32

	// (Explanation for Linux 2.6)
	// For processes running a real-time scheduling policy (policy below; see
	// sched_setscheduler(2)), this is the negated scheduling priority, minus
	// one; that is, a number in the range -2 to -100, corresponding to
	// real-time priorities 1 to 99. For processes running under a
	// non-real-time scheduling policy, this is the raw nice value
	// (setpriority(2)) as represented in the kernel. The kernel stores nice
	// values as numbers in the range 0 (high) to 39 (low), corresponding to
	// the user-visible nice range of -20 to 19.
	// Before Linux 2.6, this was a scaled value based on the scheduler
	// weighting given to this process.
	Priority int32

	// The nice value (see setpriority(2)), a value in the range 19 (low
	// priority) to -20 (high priority).
	Nice int32

	// Number of threads in this process (since Linux 2.6). Before kernel 2.6,
	// this field was hard coded to 0 as a placeholder for an earlier removed
	// field.
	NumThreads int32

	// The time in jiffies before the next SIGALRM is sent to the process due
	// to an interval timer. Since kernel 2.6.17, this field is no longer
	// maintained, and is hard coded as 0.
	Itrealvalue int32

	// The time the process started after system boot. In kernels before Linux
	// 2.6, this value was expressed in jiffies. Since Linux 2.6, the value is
	// expressed in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	// (was uint32 before Linux 2.6)
	Starttime uint64

	// Virtual memory size in bytes.
	Vsize uint32

	// Resident Set Size: number of pages the process has in real memory. This
	// is just the pages which count toward text, data, or stack space. This
	// does not include pages which have not been demand-loaded in, or which
	// are swapped out.
	Rss uint64

	// Current soft limit in bytes on the rss of the process; see the
	// description of RLIMIT_RSS in getrlimit(2).
	Rsslim uint64

	// The address above which program text can run.
	Startcode uint32

	// The address below which program text can run.
	Endcode uint32

	// The address of the start (i.e., bottom) of the stack.
	Startstack uint64

	// The current value of ESP (stack pointer), as found in the kernel stack
	// page for the process.
	Kstkesp uint64

	// The current EIP (instruction pointer).
	Kstkeip uint32

	// The bitmap of pending signals, displayed as a decimal number. Obsolete,
	// because it does not provide information on real-time signals;
	// use /proc/[pid]/status instead.
	Signal uint32

	// The bitmap of blocked signals, displayed as a decimal number. Obsolete,
	// because it does not provide information on real-time signals;
	// use /proc/[pid]/status instead.
	Blocked uint32

	// The bitmap of ignored signals, displayed as a decimal number. Obsolete,
	// because it does not provide information on real-time signals;
	// use /proc/[pid]/status instead.
	Sigignore uint32

	// The bitmap of caught signals, displayed as a decimal number. Obsolete,
	// because it does not provide information on real-time signals;
	// use /proc/[pid]/status instead.
	Sigcatch uint32

	// This is the "channel" in which the process is waiting. It is the address
	// of a location in the kernel where the process is sleeping. The
	// corresponding symbolic name can be found in /proc/[pid]/wchan.
	Wchan uint64

	// Number of pages swapped (not maintained).
	Nswap uint32

	// Cumulative nswap for child processes (not maintained).
	Cnswap uint32

	// (since Linux 2.1.22)
	// Signal to be sent to parent when we die.
	ExitSignal int

	// (since Linux 2.2.8)
	// CPU number last executed on.
	Processor int

	// (since Linux 2.5.19; was uint32 before Linux 2.6.22)
	// Real-time scheduling priority, a number in the range 1 to 99 for
	// processes scheduled under a real-time policy, or 0, for non-real-time
	// processes (see sched_setscheduler(2)).
	RtPriority uint

	// (since Linux 2.5.19; was uint32 before Linux 2.6.22)
	// Scheduling policy (see sched_setscheduler(2)). Decode using the SCHED_*
	// constants in linux/sched.h.
	Policy uint

	// (since Linux 2.6.18)
	// Aggregated block I/O delays, measured in clock ticks (centiseconds).
	DelayacctBlkioTicks uint64

	// (since Linux 2.6.24)
	// Guest time of the process (time spent running a virtual CPU for a guest
	// operating system), measured in clock ticks (divide by
	// sysconf(_SC_CLK_TCK)).
	GuestTime uint32

	// (since Linux 2.6.24)
	// Guest time of the process's children, measured in clock ticks (divide by
	// sysconf(_SC_CLK_TCK)).
	CguestTime int32
}

func NewProcPidStat(pid int) ProcPidStat {
	file := fmt.Sprintf("/proc/%d/stat", pid)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	p := ProcPidStat{}

	parsed, err := fmt.Sscanf(string(bytes),
		"%d %s %s %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d "+
			"%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d",
		&p.Pid, &p.Comm, &p.State, &p.Ppid, &p.Pgrp, &p.Session, &p.TtyNr,
		&p.Tpgid, &p.Flags, &p.Minflt, &p.Cminflt, &p.Majflt, &p.Cmajflt,
		&p.Utime, &p.Stime, &p.Cutime, &p.Cstime, &p.Priority, &p.Nice,
		&p.NumThreads, &p.Itrealvalue, &p.Starttime, &p.Vsize, &p.Rss,
		&p.Rsslim, &p.Startcode, &p.Endcode, &p.Startstack, &p.Kstkesp,
		&p.Kstkeip, &p.Signal, &p.Blocked, &p.Sigignore, &p.Sigcatch, &p.Wchan,
		&p.Nswap, &p.Cnswap, &p.ExitSignal, &p.Processor, &p.RtPriority,
		&p.Policy, &p.DelayacctBlkioTicks, &p.GuestTime, &p.CguestTime)
	if parsed < 44 {
		fmt.Println("Managed to parse only", parsed, "fields out of 44")
	}
	if err != nil {
		panic(err)
	}

	return p
}
