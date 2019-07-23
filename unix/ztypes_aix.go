// Generated code. DO NOT EDIT.

package unix

const (
	SizeofPtr              = 0x8
	SizeofShort            = 0x2
	SizeofInt              = 0x4
	SizeofLong             = 0x8
	SizeofLongLong         = 0x8
	PathMax                = 0x3ff
	SizeofSockaddrInet4    = 0x10
	SizeofSockaddrInet6    = 0x1c
	SizeofSockaddrAny      = 0x404
	SizeofSockaddrUnix     = 0x401
	SizeofSockaddrDatalink = 0x80
	SizeofLinger           = 0x8
	SizeofIPMreq           = 0x8
	SizeofIPv6Mreq         = 0x14
	SizeofIPv6MTUInfo      = 0x20
	SizeofMsghdr           = 0x30
	SizeofCmsghdr          = 0xc
	SizeofICMPv6Filter     = 0x20
	SizeofIfMsghdr         = 0x10
	AT_FDCWD               = -0x2
	AT_REMOVEDIR           = 0x1
	AT_SYMLINK_NOFOLLOW    = 0x1
	POLLERR                = 0x4000
	POLLHUP                = 0x2000
	POLLIN                 = 0x1
	POLLNVAL               = 0x8000
	POLLOUT                = 0x2
	POLLPRI                = 0x4
	POLLRDBAND             = 0x20
	POLLRDNORM             = 0x10
	POLLWRBAND             = 0x40
	POLLWRNORM             = 0x2
	RNDGETENTCNT           = 0x80045200
)

type (
	_C_short     int16
	_C_int       int32
	_C_long_long int64
	off64        int64
	Mode_t       uint32
	Timeval32    struct {
		Sec  int32
		Usec int32
	}
	Timex    struct{}
	Tms      struct{}
	Timezone struct {
		Minuteswest int32
		Dsttime     int32
	}
	Rlimit struct {
		Cur uint64
		Max uint64
	}
	Pid_t            int32
	_Gid_t           uint32
	StatxTimestamp   struct{}
	Statx_t          struct{}
	RawSockaddrInet4 struct {
		Len    uint8
		Family uint8
		Port   uint16
		Addr   [4]byte
		Zero   [8] /* in_addr */ uint8
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
		Path   [1023] /* in6_addr */ uint8
	}
	RawSockaddrDatalink struct {
		Len    uint8
		Family uint8
		Index  uint16
		Type   uint8
		Nlen   uint8
		Alen   uint8
		Slen   uint8
		Data   [120]uint8
	}
	RawSockaddr struct {
		Len    uint8
		Family uint8
		Data   [14]uint8
	}
	RawSockaddrAny struct {
		Addr RawSockaddr
		Pad  [1012]uint8
	}
	_Socklen uint32
	Cmsghdr  struct {
		Len   uint32
		Level int32
		Type  int32
	}
	ICMPv6Filter struct{ Filt [8]uint32 }
	IPMreq       struct {
		Multiaddr [4]byte
		Interface [4] /* in_addr */ byte
	}
	IPv6Mreq struct {
		Multiaddr [16] /* in_addr */ byte
		Interface uint32
	}
	IPv6MTUInfo struct {
		Addr RawSockaddrInet6
		Mtu  uint32
	}
	Linger struct {
		Onoff  int32
		Linger int32
	}
	Msghdr struct {
		Name       *byte
		Namelen    uint32
		Iov        *Iovec
		Iovlen     int32
		Control    *byte
		Controllen uint32
		Flags      int32
	}
	IfMsgHdr struct {
		Msglen  uint16
		Version uint8
		Type    uint8
		Addrs   int32
		Flags   int32
		Index   uint16
		Addrlen uint8
		_       [1] /* in6_addr */ byte
	}
	Utsname struct {
		Sysname  [32]byte
		Nodename [32]byte
		Release  [32]byte
		Version  [32]byte
		Machine  [32]byte
	}
	Ustat_t struct{}
	Termios struct {
		Iflag uint32
		Oflag uint32
		Cflag uint32
		Lflag uint32
		Cc    [16]uint8
	}
	Termio struct {
		Iflag uint16
		Oflag uint16
		Cflag uint16
		Lflag uint16
		Line  uint8
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
		Events  uint16
		Revents uint16
	}
	Flock_t struct {
		Type   int16
		Whence int16
		Sysid  uint32
		Pid    int32
		Vfs    int32
		Start  int64
		Len    int64
	}
	Fsid_t   struct{ Val [2]uint32 }
	Fsid64_t struct{ Val [2]uint64 }
)
