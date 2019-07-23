// Generated code. DO NOT EDIT.

package unix

const (
	SizeofPtr              = 0x8
	SizeofShort            = 0x2
	SizeofInt              = 0x4
	SizeofLong             = 0x8
	SizeofLongLong         = 0x8
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
	SizeofInet4Pktinfo     = 0xc
	SizeofInet6Pktinfo     = 0x14
	SizeofIPv6MTUInfo      = 0x20
	SizeofICMPv6Filter     = 0x20
	PTRACE_TRACEME         = 0x0
	PTRACE_CONT            = 0x7
	PTRACE_KILL            = 0x8
	SizeofIfMsghdr         = 0x70
	SizeofIfData           = 0x60
	SizeofIfaMsghdr        = 0x14
	SizeofIfmaMsghdr       = 0x10
	SizeofIfmaMsghdr2      = 0x14
	SizeofRtMsghdr         = 0x5c
	SizeofRtMetrics        = 0x38
	SizeofBpfVersion       = 0x4
	SizeofBpfStat          = 0x8
	SizeofBpfProgram       = 0x10
	SizeofBpfInsn          = 0x8
	SizeofBpfHdr           = 0x14
	AT_FDCWD               = -0x2
	AT_REMOVEDIR           = 0x80
	AT_SYMLINK_FOLLOW      = 0x40
	AT_SYMLINK_NOFOLLOW    = 0x20
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
	Statfs_t struct {
		Bsize       uint32
		Iosize      int32
		Blocks      uint64
		Bfree       uint64
		Bavail      uint64
		Files       uint64
		Ffree       uint64
		Fsid        Fsid
		Owner       uint32
		Type        uint32
		Flags       uint32
		Fssubtype   uint32
		Fstypename  [16]int8
		Mntonname   [1024]int8
		Mntfromname [1024]int8
		Reserved    [8]uint32
	}
	Flock_t struct {
		Start  int64
		Len    int64
		Pid    int32
		Type   int16
		Whence int16
	}
	Fstore_t struct {
		Flags      uint32
		Posmode    int32
		Offset     int64
		Length     int64
		Bytesalloc int64
	}
	Fsid   struct{ Val [2]int32 }
	Dirent struct {
		Ino     uint64
		Seekoff uint64
		Reclen  uint16
		Namlen  uint16
		Type    uint8
		Name    [1024]int8
		_       [3]byte
	}
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
	Inet4Pktinfo struct {
		Ifindex  uint32
		Spec_dst [4] /* in6_addr */ byte
		Addr     [4] /* in_addr */ byte
	}
	Inet6Pktinfo struct {
		Addr    [16] /* in_addr */ byte
		Ifindex uint32
	}
	IPv6MTUInfo struct {
		Addr RawSockaddrInet6
		Mtu  uint32
	}
	ICMPv6Filter struct {
		Filt [8] /* in6_addr */ uint32
	}
	FdSet    struct{ Bits [32]int32 }
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
	IfmaMsghdr2 struct {
		Msglen   uint16
		Version  uint8
		Type     uint8
		Addrs    int32
		Flags    int32
		Index    uint16
		_        [2]byte
		Refcount int32
	}
	RtMsghdr struct {
		Msglen  uint16
		Version uint8
		Type    uint8
		Index   uint16
		_       [2]byte
		Flags   int32
		Addrs   int32
		Pid     int32
		Seq     int32
		Errno   int32
		Use     int32
		Inits   uint32
		Rmx     RtMetrics
	}
	RtMetrics struct {
		Locks    uint32
		Mtu      uint32
		Hopcount uint32
		Expire   int32
		Recvpipe uint32
		Sendpipe uint32
		Ssthresh uint32
		Rtt      uint32
		Rttvar   uint32
		Pksent   uint32
		Filler   [4]uint32
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
