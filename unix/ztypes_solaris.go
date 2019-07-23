// Generated code. DO NOT EDIT.

package unix

const (
	SizeofPtr              = 0x8
	SizeofShort            = 0x2
	SizeofInt              = 0x4
	SizeofLong             = 0x8
	SizeofLongLong         = 0x8
	PathMax                = 0x400
	MaxHostNameLen         = 0x100
	SizeofSockaddrInet4    = 0x10
	SizeofSockaddrInet6    = 0x20
	SizeofSockaddrAny      = 0xfc
	SizeofSockaddrUnix     = 0x6e
	SizeofSockaddrDatalink = 0xfc
	SizeofLinger           = 0x8
	SizeofIPMreq           = 0x8
	SizeofIPv6Mreq         = 0x14
	SizeofMsghdr           = 0x30
	SizeofCmsghdr          = 0xc
	SizeofInet6Pktinfo     = 0x14
	SizeofIPv6MTUInfo      = 0x24
	SizeofICMPv6Filter     = 0x20
	AT_FDCWD               = 0xffd19553
	AT_SYMLINK_NOFOLLOW    = 0x1000
	AT_SYMLINK_FOLLOW      = 0x2000
	AT_REMOVEDIR           = 0x1
	AT_EACCESS             = 0x4
	SizeofIfMsghdr         = 0x54
	SizeofIfData           = 0x44
	SizeofIfaMsghdr        = 0x14
	SizeofRtMsghdr         = 0x4c
	SizeofRtMetrics        = 0x28
	SizeofBpfVersion       = 0x4
	SizeofBpfStat          = 0x80
	SizeofBpfProgram       = 0x10
	SizeofBpfInsn          = 0x8
	SizeofBpfHdr           = 0x14
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
)

type (
	_C_short     int16
	_C_int       int32
	_C_long      int64
	_C_long_long int64
	Timespec     struct {
		Sec  int64
		Nsec int64
	}
	Timeval struct {
		Sec  int64
		Usec int64
	}
	Timeval32 struct {
		Sec  int32
		Usec int32
	}
	Tms struct {
		Utime  int64
		Stime  int64
		Cutime int64
		Cstime int64
	}
	Utimbuf struct {
		Actime  int64
		Modtime int64
	}
	Rusage struct {
		Utime    Timeval
		Stime    Timeval
		Maxrss   int64
		Ixrss    int64
		Idrss    int64
		Isrss    int64
		Minflt   int64
		Majflt   int64
		Nswap    int64
		Inblock  int64
		Oublock  int64
		Msgsnd   int64
		Msgrcv   int64
		Nsignals int64
		Nvcsw    int64
		Nivcsw   int64
	}
	Rlimit struct {
		Cur uint64
		Max uint64
	}
	_Gid_t uint32
	Stat_t struct {
		Dev     uint64
		Ino     uint64
		Mode    uint32
		Nlink   uint32
		Uid     uint32
		Gid     uint32
		Rdev    uint64
		Size    int64
		Atim    Timespec
		Mtim    Timespec
		Ctim    Timespec
		Blksize int32
		_       [4]byte
		Blocks  int64
		Fstype  [16]int8
	}
	Flock_t struct {
		Type   int16
		Whence int16
		_      [4]byte
		Start  int64
		Len    int64
		Sysid  int32
		Pid    int32
		Pad    [4]int64
	}
	Dirent struct {
		Ino    uint64
		Off    int64
		Reclen uint16
		Name   [1]int8
		_      [5]byte
	}
	_Fsblkcnt_t uint64
	Statvfs_t   struct {
		Bsize    uint64
		Frsize   uint64
		Blocks   uint64
		Bfree    uint64
		Bavail   uint64
		Files    uint64
		Ffree    uint64
		Favail   uint64
		Fsid     uint64
		Basetype [16]int8
		Flag     uint64
		Namemax  uint64
		Fstr     [32]int8
	}
	RawSockaddrInet4 struct {
		Family uint16
		Port   uint16
		Addr   [4]byte
		Zero   [8] /* in_addr */ int8
	}
	RawSockaddrInet6 struct {
		Family         uint16
		Port           uint16
		Flowinfo       uint32
		Addr           [16]byte
		Scope_id       uint32
		X__sin6_src_id uint32
	}
	RawSockaddrUnix struct {
		Family uint16
		Path   [108] /* in6_addr */ int8
	}
	RawSockaddrDatalink struct {
		Family uint16
		Index  uint16
		Type   uint8
		Nlen   uint8
		Alen   uint8
		Slen   uint8
		Data   [244]int8
	}
	RawSockaddr struct {
		Family uint16
		Data   [14]int8
	}
	RawSockaddrAny struct {
		Addr RawSockaddr
		Pad  [236]int8
	}
	_Socklen uint32
	Linger   struct {
		Onoff  int32
		Linger int32
	}
	Iovec struct {
		Base *int8
		Len  uint64
	}
	IPMreq struct {
		Multiaddr [4]byte
		Interface [4] /* in_addr */ byte
	}
	IPv6Mreq struct {
		Multiaddr [16] /* in_addr */ byte
		Interface uint32
	}
	Msghdr struct {
		Name         *byte
		Namelen      uint32
		_            [4] /* in6_addr */ byte
		Iov          *Iovec
		Iovlen       int32
		_            [4]byte
		Accrights    *int8
		Accrightslen int32
		_            [4]byte
	}
	Cmsghdr struct {
		Len   uint32
		Level int32
		Type  int32
	}
	Inet6Pktinfo struct {
		Addr    [16]byte
		Ifindex uint32
	}
	IPv6MTUInfo struct {
		Addr RawSockaddrInet6
		Mtu  uint32
	}
	ICMPv6Filter struct {
		X__icmp6_filt [8] /* in6_addr */ uint32
	}
	FdSet   struct{ Bits [1024]int64 }
	Utsname struct {
		Sysname  [257]byte
		Nodename [257]byte
		Release  [257]byte
		Version  [257]byte
		Machine  [257]byte
	}
	Ustat_t struct {
		Tfree  int64
		Tinode uint64
		Fname  [6]int8
		Fpack  [6]int8
		_      [4]byte
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
	IfData struct {
		Type       uint8
		Addrlen    uint8
		Hdrlen     uint8
		_          [1]byte
		Mtu        uint32
		Metric     uint32
		Baudrate   uint32
		Ipackets   uint32
		Ierrors    uint32
		Opackets   uint32
		Oerrors    uint32
		Collisions uint32
		Ibytes     uint32
		Obytes     uint32
		Imcasts    uint32
		Omcasts    uint32
		Iqdrops    uint32
		Noproto    uint32
		Lastchange Timeval32
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
		Expire   uint32
		Recvpipe uint32
		Sendpipe uint32
		Ssthresh uint32
		Rtt      uint32
		Rttvar   uint32
		Pksent   uint32
	}
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
	BpfProgram struct {
		Len   uint32
		_     [4]byte
		Insns *BpfInsn
	}
	BpfInsn struct {
		Code uint16
		Jt   uint8
		Jf   uint8
		K    uint32
	}
	BpfTimeval struct {
		Sec  int32
		Usec int32
	}
	BpfHdr struct {
		Tstamp  BpfTimeval
		Caplen  uint32
		Datalen uint32
		Hdrlen  uint16
		_       [2]byte
	}
	Termios struct {
		Iflag uint32
		Oflag uint32
		Cflag uint32
		Lflag uint32
		Cc    [19]uint8
		_     [1]byte
	}
	Termio struct {
		Iflag uint16
		Oflag uint16
		Cflag uint16
		Lflag uint16
		Line  int8
		Cc    [8]uint8
		_     [1]byte
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
)
