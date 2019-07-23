// Generated code. DO NOT EDIT.

package unix

const (
	SizeofPtr              = 0x8
	SizeofShort            = 0x2
	SizeofInt              = 0x4
	SizeofLong             = 0x8
	SizeofLongLong         = 0x8
	_statfsVersion         = 0x20140518
	_dirblksiz             = 0x400
	PathMax                = 0x400
	FADV_NORMAL            = 0x0
	FADV_RANDOM            = 0x1
	FADV_SEQUENTIAL        = 0x2
	FADV_WILLNEED          = 0x3
	FADV_DONTNEED          = 0x4
	FADV_NOREUSE           = 0x5
	SizeofSockaddrInet4    = 0x10
	SizeofSockaddrInet6    = 0x1c
	SizeofSockaddrAny      = 0x6c
	SizeofSockaddrUnix     = 0x6a
	SizeofSockaddrDatalink = 0x36
	SizeofLinger           = 0x8
	SizeofIPMreq           = 0x8
	SizeofIPMreqn          = 0xc
	SizeofIPv6Mreq         = 0x14
	SizeofMsghdr           = 0x30
	SizeofCmsghdr          = 0xc
	SizeofInet6Pktinfo     = 0x14
	SizeofIPv6MTUInfo      = 0x20
	SizeofICMPv6Filter     = 0x20
	PTRACE_ATTACH          = 0xa
	PTRACE_CONT            = 0x7
	PTRACE_DETACH          = 0xb
	PTRACE_GETFPREGS       = 0x23
	PTRACE_GETFSBASE       = 0x47
	PTRACE_GETLWPLIST      = 0xf
	PTRACE_GETNUMLWPS      = 0xe
	PTRACE_GETREGS         = 0x21
	PTRACE_GETXSTATE       = 0x45
	PTRACE_IO              = 0xc
	PTRACE_KILL            = 0x8
	PTRACE_LWPEVENTS       = 0x18
	PTRACE_LWPINFO         = 0xd
	PTRACE_SETFPREGS       = 0x24
	PTRACE_SETREGS         = 0x22
	PTRACE_SINGLESTEP      = 0x9
	PTRACE_TRACEME         = 0x0
	PIOD_READ_D            = 0x1
	PIOD_WRITE_D           = 0x2
	PIOD_READ_I            = 0x3
	PIOD_WRITE_I           = 0x4
	PL_FLAG_BORN           = 0x100
	PL_FLAG_EXITED         = 0x200
	PL_FLAG_SI             = 0x20
	TRAP_BRKPT             = 0x1
	TRAP_TRACE             = 0x2
	sizeofIfMsghdr         = 0xa8
	SizeofIfMsghdr         = 0xa8
	sizeofIfData           = 0x98
	SizeofIfData           = 0x98
	SizeofIfaMsghdr        = 0x14
	SizeofIfmaMsghdr       = 0x10
	SizeofIfAnnounceMsghdr = 0x18
	SizeofRtMsghdr         = 0x98
	SizeofRtMetrics        = 0x70
	SizeofBpfVersion       = 0x4
	SizeofBpfStat          = 0x8
	SizeofBpfZbuf          = 0x18
	SizeofBpfProgram       = 0x10
	SizeofBpfInsn          = 0x8
	SizeofBpfHdr           = 0x20
	SizeofBpfZbufHeader    = 0x20
	AT_FDCWD               = -0x64
	AT_REMOVEDIR           = 0x800
	AT_SYMLINK_FOLLOW      = 0x400
	AT_SYMLINK_NOFOLLOW    = 0x200
	POLLERR                = 0x8
	POLLHUP                = 0x10
	POLLIN                 = 0x1
	POLLINIGNEOF           = 0x2000
	POLLNVAL               = 0x20
	POLLOUT                = 0x4
	POLLPRI                = 0x2
	POLLRDBAND             = 0x80
	POLLRDNORM             = 0x40
	POLLWRBAND             = 0x100
	POLLWRNORM             = 0x4
)

type (
	_C_short     int16
	_C_int       int32
	_C_long_long int64
	Rlimit       struct {
		Cur int64
		Max int64
	}
	_Gid_t   uint32
	Statfs_t struct {
		Version     uint32
		Type        uint32
		Flags       uint64
		Bsize       uint64
		Iosize      uint64
		Blocks      uint64
		Bfree       uint64
		Bavail      int64
		Files       uint64
		Ffree       int64
		Syncwrites  uint64
		Asyncwrites uint64
		Syncreads   uint64
		Asyncreads  uint64
		Spare       [10]uint64
		Namemax     uint32
		Owner       uint32
		Fsid        Fsid
		Charspare   [80]int8
		Fstypename  [16]int8
		Mntfromname [1024]int8
		Mntonname   [1024]int8
	}
	statfs_freebsd11_t struct {
		Version     uint32
		Type        uint32
		Flags       uint64
		Bsize       uint64
		Iosize      uint64
		Blocks      uint64
		Bfree       uint64
		Bavail      int64
		Files       uint64
		Ffree       int64
		Syncwrites  uint64
		Asyncwrites uint64
		Syncreads   uint64
		Asyncreads  uint64
		Spare       [10]uint64
		Namemax     uint32
		Owner       uint32
		Fsid        Fsid
		Charspare   [80]int8
		Fstypename  [16]int8
		Mntfromname [88]int8
		Mntonname   [88]int8
	}
	Dirent struct {
		Fileno uint64
		Off    int64
		Reclen uint16
		Type   uint8
		Pad0   uint8
		Namlen uint16
		Pad1   uint16
		Name   [256]int8
	}
	dirent_freebsd11 struct {
		Fileno uint32
		Reclen uint16
		Type   uint8
		Namlen uint8
		Name   [256]int8
	}
	Fsid             struct{ Val [2]int32 }
	RawSockaddrInet4 struct {
		Len    uint8
		Family uint8
		Port   uint16
		Addr   [4]byte
		Zero   [8] /* in_addr */ int8
	}
	RawSockaddrInet6 struct {
		Len      uint8
		Family   uint8
		Port     uint16
		Flowinfo uint32
		Addr     [16]byte
		Scope_id uint32
	}
	RawSockaddrUnix struct {
		Len    uint8
		Family uint8
		Path   [104] /* in6_addr */ int8
	}
	RawSockaddrDatalink struct {
		Len    uint8
		Family uint8
		Index  uint16
		Type   uint8
		Nlen   uint8
		Alen   uint8
		Slen   uint8
		Data   [46]int8
	}
	RawSockaddr struct {
		Len    uint8
		Family uint8
		Data   [14]int8
	}
	RawSockaddrAny struct {
		Addr RawSockaddr
		Pad  [92]int8
	}
	_Socklen uint32
	Linger   struct {
		Onoff  int32
		Linger int32
	}
	IPMreq struct {
		Multiaddr [4]byte
		Interface [4] /* in_addr */ byte
	}
	IPMreqn struct {
		Multiaddr [4] /* in_addr */ byte
		Address   [4] /* in_addr */ byte
		Ifindex   int32
	}
	IPv6Mreq struct {
		Multiaddr [16] /* in_addr */ byte
		Interface uint32
	}
	Cmsghdr struct {
		Len   uint32
		Level int32
		Type  int32
	}
	Inet6Pktinfo struct {
		Addr    [16] /* in6_addr */ byte
		Ifindex uint32
	}
	IPv6MTUInfo struct {
		Addr RawSockaddrInet6
		Mtu  uint32
	}
	ICMPv6Filter struct {
		Filt [8] /* in6_addr */ uint32
	}
	PtraceLwpInfoStruct struct {
		Lwpid        int32
		Event        int32
		Flags        int32
		Sigmask      Sigset_t
		Siglist      Sigset_t
		Siginfo      __Siginfo
		Tdname       [20]int8
		Child_pid    int32
		Syscall_code uint32
		Syscall_narg uint32
	}
	Sigset_t     struct{ Val [4]uint32 }
	PtraceIoDesc struct {
		Op   int32
		Offs *byte
		Addr *byte
		Len  uint
	}
	ifMsghdr struct {
		Msglen  uint16
		Version uint8
		Type    uint8
		Addrs   int32
		Flags   int32
		Index   uint16
		_       [2]byte
		Data    ifData
	}
	IfMsghdr struct {
		Msglen  uint16
		Version uint8
		Type    uint8
		Addrs   int32
		Flags   int32
		Index   uint16
		_       [2]byte
		Data    IfData
	}
	ifData struct {
		Type       uint8
		Physical   uint8
		Addrlen    uint8
		Hdrlen     uint8
		Link_state uint8
		Vhid       uint8
		Datalen    uint16
		Mtu        uint32
		Metric     uint32
		Baudrate   uint64
		Ipackets   uint64
		Ierrors    uint64
		Opackets   uint64
		Oerrors    uint64
		Collisions uint64
		Ibytes     uint64
		Obytes     uint64
		Imcasts    uint64
		Omcasts    uint64
		Iqdrops    uint64
		Oqdrops    uint64
		Noproto    uint64
		Hwassist   uint64
		_          [8]byte
		_          [16]byte
	}
	IfaMsghdr struct {
		Msglen  uint16
		Version uint8
		Type    uint8
		Addrs   int32
		Flags   int32
		Index   uint16
		_       [2]byte
		Metric  int32
	}
	IfmaMsghdr struct {
		Msglen  uint16
		Version uint8
		Type    uint8
		Addrs   int32
		Flags   int32
		Index   uint16
		_       [2]byte
	}
	IfAnnounceMsghdr struct {
		Msglen  uint16
		Version uint8
		Type    uint8
		Index   uint16
		Name    [16]int8
		What    uint16
	}
	BpfVersion struct {
		Major uint16
		Minor uint16
	}
	BpfStat struct {
		Recv uint32
		Drop uint32
	}
	BpfInsn struct {
		Code uint16
		Jt   uint8
		Jf   uint8
		K    uint32
	}
	BpfZbufHeader struct {
		Kernel_gen uint32
		Kernel_len uint32
		User_gen   uint32
		_          [5]uint32
	}
	Termios struct {
		Iflag  uint32
		Oflag  uint32
		Cflag  uint32
		Lflag  uint32
		Cc     [20]uint8
		Ispeed uint32
		Ospeed uint32
	}
	Winsize struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}
	PollFd struct {
		Fd      int32
		Events  int16
		Revents int16
	}
	CapRights struct{ Rights [2]uint64 }
	Utsname   struct {
		Sysname  [256]byte
		Nodename [256]byte
		Release  [256]byte
		Version  [256]byte
		Machine  [256]byte
	}
)
