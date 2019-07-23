// Generated code. DO NOT EDIT.

package unix

const (
	SizeofPtr              = 0x8
	SizeofShort            = 0x2
	SizeofInt              = 0x4
	SizeofLong             = 0x8
	SizeofLongLong         = 0x8
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
	SizeofSockaddrDatalink = 0x14
	SizeofLinger           = 0x8
	SizeofIPMreq           = 0x8
	SizeofIPv6Mreq         = 0x14
	SizeofMsghdr           = 0x30
	SizeofCmsghdr          = 0xc
	SizeofInet6Pktinfo     = 0x14
	SizeofIPv6MTUInfo      = 0x20
	SizeofICMPv6Filter     = 0x20
	PTRACE_TRACEME         = 0x0
	PTRACE_CONT            = 0x7
	PTRACE_KILL            = 0x8
	SizeofIfMsghdr         = 0x98
	SizeofIfData           = 0x88
	SizeofIfaMsghdr        = 0x18
	SizeofIfAnnounceMsghdr = 0x18
	SizeofRtMsghdr         = 0x78
	SizeofRtMetrics        = 0x50
	SizeofBpfVersion       = 0x4
	SizeofBpfStat          = 0x80
	SizeofBpfProgram       = 0x10
	SizeofBpfInsn          = 0x8
	SizeofBpfHdr           = 0x20
	AT_FDCWD               = -0x64
	AT_SYMLINK_FOLLOW      = 0x400
	AT_SYMLINK_NOFOLLOW    = 0x200
	POLLERR                = 0x8
	POLLHUP                = 0x10
	POLLIN                 = 0x1
	POLLNVAL               = 0x20
	POLLOUT                = 0x4
	POLLPRI                = 0x2
	POLLRDBAND             = 0x80
	POLLRDNORM             = 0x40
	POLLWRBAND             = 0x100
	POLLWRNORM             = 0x4
	SizeofClockinfo        = 0x14
)

type (
	_C_short     int16
	_C_int       int32
	_C_long_long int64
	Rlimit       struct {
		Cur uint64
		Max uint64
	}
	_Gid_t   uint32
	Statfs_t [0]byte
	Flock_t  struct {
		Start  int64
		Len    int64
		Pid    int32
		Type   int16
		Whence int16
	}
	Dirent struct {
		Fileno    uint64
		Reclen    uint16
		Namlen    uint16
		Type      uint8
		Name      [512]int8
		Pad_cgo_0 [3]byte
	}
	Fsid             struct{ X__fsid_val [2]int32 }
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
		Data   [12]int8
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
	FdSet  struct{ Bits [8]uint32 }
	IfData struct {
		Type       uint8
		Addrlen    uint8
		Hdrlen     uint8
		Pad_cgo_0  [1]byte
		Link_state int32
		Mtu        uint64
		Metric     uint64
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
		Noproto    uint64
		Lastchange Timespec
	}
	IfaMsghdr struct {
		Msglen    uint16
		Version   uint8
		Type      uint8
		Addrs     int32
		Flags     int32
		Metric    int32
		Index     uint16
		Pad_cgo_0 [6]byte
	}
	IfAnnounceMsghdr struct {
		Msglen  uint16
		Version uint8
		Type    uint8
		Index   uint16
		Name    [16]int8
		What    uint16
	}
	RtMsghdr struct {
		Msglen    uint16
		Version   uint8
		Type      uint8
		Index     uint16
		Pad_cgo_0 [2]byte
		Flags     int32
		Addrs     int32
		Pid       int32
		Seq       int32
		Errno     int32
		Use       int32
		Inits     int32
		Pad_cgo_1 [4]byte
		Rmx       RtMetrics
	}
	RtMetrics struct {
		Locks    uint64
		Mtu      uint64
		Hopcount uint64
		Recvpipe uint64
		Sendpipe uint64
		Ssthresh uint64
		Rtt      uint64
		Rttvar   uint64
		Expire   int64
		Pksent   int64
	}
	Mclpool    [0]byte
	BpfVersion struct {
		Major uint16
		Minor uint16
	}
	BpfStat struct {
		Recv    uint64
		Drop    uint64
		Capt    uint64
		Padding [13]uint64
	}
	BpfInsn struct {
		Code uint16
		Jt   uint8
		Jf   uint8
		K    uint32
	}
	Termios struct {
		Iflag  uint32
		Oflag  uint32
		Cflag  uint32
		Lflag  uint32
		Cc     [20]uint8
		Ispeed int32
		Ospeed int32
	}
	Winsize struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}
	Ptmget struct {
		Cfd int32
		Sfd int32
		Cn  [1024]byte
		Sn  [1024]byte
	}
	PollFd struct {
		Fd      int32
		Events  int16
		Revents int16
	}
	Sysctlnode struct {
		Flags           uint32
		Num             int32
		Name            [32]int8
		Ver             uint32
		X__rsvd         uint32
		Un              [16]byte
		X_sysctl_size   [8]byte
		X_sysctl_func   [8]byte
		X_sysctl_parent [8]byte
		X_sysctl_desc   [8]byte
	}
	Utsname struct {
		Sysname  [256]byte
		Nodename [256]byte
		Release  [256]byte
		Version  [256]byte
		Machine  [256]byte
	}
	Clockinfo struct {
		Hz      int32
		Tick    int32
		Tickadj int32
		Stathz  int32
		Profhz  int32
	}
)
