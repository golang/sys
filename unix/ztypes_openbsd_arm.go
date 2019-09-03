// cgo -godefs -- -fsigned-char types_openbsd.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build arm,openbsd

package unix

const (
	SizeofPtr = 0x4

	SizeofLong = 0x4
)

type (
	_C_long int32
)

type Timespec struct {
	Sec  int64
	Nsec int32
	_    [4]byte
}

type Timeval struct {
	Sec  int64
	Usec int32
	_    [4]byte
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int32
	Ixrss    int32
	Idrss    int32
	Isrss    int32
	Minflt   int32
	Majflt   int32
	Nswap    int32
	Inblock  int32
	Oublock  int32
	Msgsnd   int32
	Msgrcv   int32
	Nsignals int32
	Nvcsw    int32
	Nivcsw   int32
}

type Stat_t struct {
	Mode    uint32
	Dev     int32
	Ino     uint64
	Nlink   uint32
	Uid     uint32
	Gid     uint32
	Rdev    int32
	Atim    Timespec
	Mtim    Timespec
	Ctim    Timespec
	Size    int64
	Blocks  int64
	Blksize int32
	Flags   uint32
	Gen     uint32
	_       [4]byte
	_       Timespec
}

type Statfs_t struct {
	F_flags       uint32
	F_bsize       uint32
	F_iosize      uint32
	_             [4]byte
	F_blocks      uint64
	F_bfree       uint64
	F_bavail      int64
	F_files       uint64
	F_ffree       uint64
	F_favail      int64
	F_syncwrites  uint64
	F_syncreads   uint64
	F_asyncwrites uint64
	F_asyncreads  uint64
	F_fsid        Fsid
	F_namemax     uint32
	F_owner       uint32
	F_ctime       uint64
	F_fstypename  [16]int8
	F_mntonname   [90]int8
	F_mntfromname [90]int8
	F_mntfromspec [90]int8
	_             [2]byte
	Mount_info    [160]byte
}

type Dirent struct {
	Fileno uint64
	Off    int64
	Reclen uint16
	Type   uint8
	Namlen uint8
	_      [4]uint8
	Name   [256]int8
}

/* in_addr */

/* in6_addr */

type Iovec struct {
	Base *byte
	Len  uint32
}

/* in_addr */
/* in_addr */

/* in6_addr */

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     uint32
	Control    *byte
	Controllen uint32
	Flags      int32
}

/* in6_addr */

const (
	SizeofMsghdr = 0x1c
)

type Kevent_t struct {
	Ident  uint32
	Filter int16
	Flags  uint16
	Fflags uint32
	_      [4]byte
	Data   int64
	Udata  *byte
	_      [4]byte
}

const (
	SizeofIfMsghdr = 0xa8
	SizeofIfData   = 0x90
)

type IfData struct {
	Type         uint8
	Addrlen      uint8
	Hdrlen       uint8
	Link_state   uint8
	Mtu          uint32
	Metric       uint32
	Rdomain      uint32
	Baudrate     uint64
	Ipackets     uint64
	Ierrors      uint64
	Opackets     uint64
	Oerrors      uint64
	Collisions   uint64
	Ibytes       uint64
	Obytes       uint64
	Imcasts      uint64
	Omcasts      uint64
	Iqdrops      uint64
	Oqdrops      uint64
	Noproto      uint64
	Capabilities uint32
	_            [4]byte
	Lastchange   Timeval
}

type Mclpool struct{}

const (
	SizeofBpfProgram = 0x8
)

type BpfProgram struct {
	Len   uint32
	Insns *BpfInsn
}

type BpfHdr struct {
	Tstamp  BpfTimeval
	Caplen  uint32
	Datalen uint32
	Hdrlen  uint16
	_       [2]byte
}

type Uvmexp struct {
	Pagesize           int32
	Pagemask           int32
	Pageshift          int32
	Npages             int32
	Free               int32
	Active             int32
	Inactive           int32
	Paging             int32
	Wired              int32
	Zeropages          int32
	Reserve_pagedaemon int32
	Reserve_kernel     int32
	Unused01           int32
	Vnodepages         int32
	Vtextpages         int32
	Freemin            int32
	Freetarg           int32
	Inactarg           int32
	Wiredmax           int32
	Anonmin            int32
	Vtextmin           int32
	Vnodemin           int32
	Anonminpct         int32
	Vtextminpct        int32
	Vnodeminpct        int32
	Nswapdev           int32
	Swpages            int32
	Swpginuse          int32
	Swpgonly           int32
	Nswget             int32
	Nanon              int32
	Unused05           int32
	Unused06           int32
	Faults             int32
	Traps              int32
	Intrs              int32
	Swtch              int32
	Softs              int32
	Syscalls           int32
	Pageins            int32
	Unused07           int32
	Unused08           int32
	Pgswapin           int32
	Pgswapout          int32
	Forks              int32
	Forks_ppwait       int32
	Forks_sharevm      int32
	Pga_zerohit        int32
	Pga_zeromiss       int32
	Unused09           int32
	Fltnoram           int32
	Fltnoanon          int32
	Fltnoamap          int32
	Fltpgwait          int32
	Fltpgrele          int32
	Fltrelck           int32
	Fltrelckok         int32
	Fltanget           int32
	Fltanretry         int32
	Fltamcopy          int32
	Fltnamap           int32
	Fltnomap           int32
	Fltlget            int32
	Fltget             int32
	Flt_anon           int32
	Flt_acow           int32
	Flt_obj            int32
	Flt_prcopy         int32
	Flt_przero         int32
	Pdwoke             int32
	Pdrevs             int32
	Pdswout            int32
	Pdfreed            int32
	Pdscans            int32
	Pdanscan           int32
	Pdobscan           int32
	Pdreact            int32
	Pdbusy             int32
	Pdpageouts         int32
	Pdpending          int32
	Pddeact            int32
	Unused11           int32
	Unused12           int32
	Unused13           int32
	Fpswtch            int32
	Kmapent            int32
}
