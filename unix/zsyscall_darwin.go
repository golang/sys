// Generated code. DO NOT EDIT.

package unix

import "unsafe"

func getgroups(ngid int, gid *_Gid_t) (n int, err error) {
	r0, _, e1 := syscall_rawSyscall(funcPC(libc_getgroups_trampoline), uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getgroups_trampoline()
func setgroups(ngid int, gid *_Gid_t) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_setgroups_trampoline), uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setgroups_trampoline()
func wait4(pid int, wstatus *_C_int, options int, rusage *Rusage) (wpid int, err error) {
	r0, _, e1 := syscall_syscall6(funcPC(libc_wait4_trampoline), uintptr(pid), uintptr(unsafe.Pointer(wstatus)), uintptr(options), uintptr(unsafe.Pointer(rusage)), 0, 0)
	wpid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_wait4_trampoline()
func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_accept_trampoline), uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_accept_trampoline()
func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_bind_trampoline), uintptr(s), uintptr(addr), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_bind_trampoline()
func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_connect_trampoline), uintptr(s), uintptr(addr), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_connect_trampoline()
func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, _, e1 := syscall_rawSyscall(funcPC(libc_socket_trampoline), uintptr(domain), uintptr(typ), uintptr(proto))
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_socket_trampoline()
func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, _, e1 := syscall_syscall6(funcPC(libc_getsockopt_trampoline), uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getsockopt_trampoline()
func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	_, _, e1 := syscall_syscall6(funcPC(libc_setsockopt_trampoline), uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setsockopt_trampoline()
func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_getpeername_trampoline), uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getpeername_trampoline()
func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_getsockname_trampoline), uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getsockname_trampoline()
func Shutdown(s int, how int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_shutdown_trampoline), uintptr(s), uintptr(how), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_shutdown_trampoline()
func socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) {
	_, _, e1 := syscall_rawSyscall6(funcPC(libc_socketpair_trampoline), uintptr(domain), uintptr(typ), uintptr(proto), uintptr(unsafe.Pointer(fd)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_socketpair_trampoline()
func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := syscall_syscall6(funcPC(libc_recvfrom_trampoline), uintptr(fd), uintptr(_p0), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_recvfrom_trampoline()
func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 unsafe.Pointer
	if len(buf) > 0 {
		_p0 = unsafe.Pointer(&buf[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := syscall_syscall6(funcPC(libc_sendto_trampoline), uintptr(s), uintptr(_p0), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_sendto_trampoline()
func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_recvmsg_trampoline), uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_recvmsg_trampoline()
func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_sendmsg_trampoline), uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_sendmsg_trampoline()
func kevent(kq int, change unsafe.Pointer, nchange int, event unsafe.Pointer, nevent int, timeout *Timespec) (n int, err error) {
	r0, _, e1 := syscall_syscall6(funcPC(libc_kevent_trampoline), uintptr(kq), uintptr(change), uintptr(nchange), uintptr(event), uintptr(nevent), uintptr(unsafe.Pointer(timeout)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_kevent_trampoline()
func sysctl(mib []_C_int, old *byte, oldlen *uintptr, new *byte, newlen uintptr) (err error) {
	var _p0 unsafe.Pointer
	if len(mib) > 0 {
		_p0 = unsafe.Pointer(&mib[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := syscall_syscall6(funcPC(libc___sysctl_trampoline), uintptr(_p0), uintptr(len(mib)), uintptr(unsafe.Pointer(old)), uintptr(unsafe.Pointer(oldlen)), uintptr(unsafe.Pointer(new)), uintptr(newlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc___sysctl_trampoline()
func utimes(path string, timeval *[2]Timeval) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_utimes_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(timeval)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_utimes_trampoline()
func futimes(fd int, timeval *[2]Timeval) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_futimes_trampoline), uintptr(fd), uintptr(unsafe.Pointer(timeval)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_futimes_trampoline()
func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_fcntl_trampoline), uintptr(fd), uintptr(cmd), uintptr(arg))
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fcntl_trampoline()
func poll(fds *PollFd, nfds int, timeout int) (n int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_poll_trampoline), uintptr(unsafe.Pointer(fds)), uintptr(nfds), uintptr(timeout))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_poll_trampoline()
func Madvise(b []byte, behav int) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := syscall_syscall(funcPC(libc_madvise_trampoline), uintptr(_p0), uintptr(len(b)), uintptr(behav))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_madvise_trampoline()
func Mlock(b []byte) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := syscall_syscall(funcPC(libc_mlock_trampoline), uintptr(_p0), uintptr(len(b)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_mlock_trampoline()
func Mlockall(flags int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_mlockall_trampoline), uintptr(flags), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_mlockall_trampoline()
func Mprotect(b []byte, prot int) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := syscall_syscall(funcPC(libc_mprotect_trampoline), uintptr(_p0), uintptr(len(b)), uintptr(prot))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_mprotect_trampoline()
func Msync(b []byte, flags int) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := syscall_syscall(funcPC(libc_msync_trampoline), uintptr(_p0), uintptr(len(b)), uintptr(flags))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_msync_trampoline()
func Munlock(b []byte) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := syscall_syscall(funcPC(libc_munlock_trampoline), uintptr(_p0), uintptr(len(b)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_munlock_trampoline()
func Munlockall() (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_munlockall_trampoline), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_munlockall_trampoline()
func ptrace(request int, pid int, addr uintptr, data uintptr) (err error) {
	_, _, e1 := syscall_syscall6(funcPC(libc_ptrace_trampoline), uintptr(request), uintptr(pid), uintptr(addr), uintptr(data), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_ptrace_trampoline()
func getattrlist(path *byte, list unsafe.Pointer, buf unsafe.Pointer, size uintptr, options int) (err error) {
	_, _, e1 := syscall_syscall6(funcPC(libc_getattrlist_trampoline), uintptr(unsafe.Pointer(path)), uintptr(list), uintptr(buf), uintptr(size), uintptr(options), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getattrlist_trampoline()
func pipe() (r int, w int, err error) {
	r0, r1, e1 := syscall_rawSyscall(funcPC(libc_pipe_trampoline), 0, 0, 0)
	r = int(r0)
	w = int(r1)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_pipe_trampoline()
func getxattr(path string, attr string, dest *byte, size int, position uint32, options int) (sz int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(attr)
	if err != nil {
		return
	}
	r0, _, e1 := syscall_syscall6(funcPC(libc_getxattr_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(unsafe.Pointer(dest)), uintptr(size), uintptr(position), uintptr(options))
	sz = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getxattr_trampoline()
func fgetxattr(fd int, attr string, dest *byte, size int, position uint32, options int) (sz int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(attr)
	if err != nil {
		return
	}
	r0, _, e1 := syscall_syscall6(funcPC(libc_fgetxattr_trampoline), uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(dest)), uintptr(size), uintptr(position), uintptr(options))
	sz = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fgetxattr_trampoline()
func setxattr(path string, attr string, data *byte, size int, position uint32, options int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(attr)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall6(funcPC(libc_setxattr_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(unsafe.Pointer(data)), uintptr(size), uintptr(position), uintptr(options))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setxattr_trampoline()
func fsetxattr(fd int, attr string, data *byte, size int, position uint32, options int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(attr)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall6(funcPC(libc_fsetxattr_trampoline), uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(data)), uintptr(size), uintptr(position), uintptr(options))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fsetxattr_trampoline()
func removexattr(path string, attr string, options int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(attr)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_removexattr_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(options))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_removexattr_trampoline()
func fremovexattr(fd int, attr string, options int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(attr)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_fremovexattr_trampoline), uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(options))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fremovexattr_trampoline()
func listxattr(path string, dest *byte, size int, options int) (sz int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := syscall_syscall6(funcPC(libc_listxattr_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(dest)), uintptr(size), uintptr(options), 0, 0)
	sz = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_listxattr_trampoline()
func flistxattr(fd int, dest *byte, size int, options int) (sz int, err error) {
	r0, _, e1 := syscall_syscall6(funcPC(libc_flistxattr_trampoline), uintptr(fd), uintptr(unsafe.Pointer(dest)), uintptr(size), uintptr(options), 0, 0)
	sz = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_flistxattr_trampoline()
func setattrlist(path *byte, list unsafe.Pointer, buf unsafe.Pointer, size uintptr, options int) (err error) {
	_, _, e1 := syscall_syscall6(funcPC(libc_setattrlist_trampoline), uintptr(unsafe.Pointer(path)), uintptr(list), uintptr(buf), uintptr(size), uintptr(options), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setattrlist_trampoline()
func kill(pid int, signum int, posix int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_kill_trampoline), uintptr(pid), uintptr(signum), uintptr(posix))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_kill_trampoline()
func ioctl(fd int, req uint, arg uintptr) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_ioctl_trampoline), uintptr(fd), uintptr(req), uintptr(arg))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_ioctl_trampoline()
func libc_sendfile_trampoline()
func Access(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_access_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_access_trampoline()
func Adjtime(delta *Timeval, olddelta *Timeval) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_adjtime_trampoline), uintptr(unsafe.Pointer(delta)), uintptr(unsafe.Pointer(olddelta)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_adjtime_trampoline()
func Chdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_chdir_trampoline), uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_chdir_trampoline()
func Chflags(path string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_chflags_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(flags), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_chflags_trampoline()
func Chmod(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_chmod_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_chmod_trampoline()
func Chown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_chown_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_chown_trampoline()
func Chroot(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_chroot_trampoline), uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_chroot_trampoline()
func Close(fd int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_close_trampoline), uintptr(fd), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_close_trampoline()
func Dup(fd int) (nfd int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_dup_trampoline), uintptr(fd), 0, 0)
	nfd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_dup_trampoline()
func Dup2(from int, to int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_dup2_trampoline), uintptr(from), uintptr(to), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_dup2_trampoline()
func Exchangedata(path1 string, path2 string, options int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path1)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(path2)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_exchangedata_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(options))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_exchangedata_trampoline()
func Exit(code int) {
	syscall_syscall(funcPC(libc_exit_trampoline), uintptr(code), 0, 0)
	return
}
func libc_exit_trampoline()
func Faccessat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall6(funcPC(libc_faccessat_trampoline), uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(flags), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_faccessat_trampoline()
func Fchdir(fd int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_fchdir_trampoline), uintptr(fd), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fchdir_trampoline()
func Fchflags(fd int, flags int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_fchflags_trampoline), uintptr(fd), uintptr(flags), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fchflags_trampoline()
func Fchmod(fd int, mode uint32) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_fchmod_trampoline), uintptr(fd), uintptr(mode), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fchmod_trampoline()
func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall6(funcPC(libc_fchmodat_trampoline), uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(flags), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fchmodat_trampoline()
func Fchown(fd int, uid int, gid int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_fchown_trampoline), uintptr(fd), uintptr(uid), uintptr(gid))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fchown_trampoline()
func Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall6(funcPC(libc_fchownat_trampoline), uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), uintptr(flags), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fchownat_trampoline()
func Flock(fd int, how int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_flock_trampoline), uintptr(fd), uintptr(how), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_flock_trampoline()
func Fpathconf(fd int, name int) (val int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_fpathconf_trampoline), uintptr(fd), uintptr(name), 0)
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fpathconf_trampoline()
func Fsync(fd int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_fsync_trampoline), uintptr(fd), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_fsync_trampoline()
func libc_ftruncate_trampoline()
func Getdtablesize() (size int) {
	r0, _, _ := syscall_syscall(funcPC(libc_getdtablesize_trampoline), 0, 0, 0)
	size = int(r0)
	return
}
func libc_getdtablesize_trampoline()
func Getegid() (egid int) {
	r0, _, _ := syscall_rawSyscall(funcPC(libc_getegid_trampoline), 0, 0, 0)
	egid = int(r0)
	return
}
func libc_getegid_trampoline()
func Geteuid() (uid int) {
	r0, _, _ := syscall_rawSyscall(funcPC(libc_geteuid_trampoline), 0, 0, 0)
	uid = int(r0)
	return
}
func libc_geteuid_trampoline()
func Getgid() (gid int) {
	r0, _, _ := syscall_rawSyscall(funcPC(libc_getgid_trampoline), 0, 0, 0)
	gid = int(r0)
	return
}
func libc_getgid_trampoline()
func Getpgid(pid int) (pgid int, err error) {
	r0, _, e1 := syscall_rawSyscall(funcPC(libc_getpgid_trampoline), uintptr(pid), 0, 0)
	pgid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getpgid_trampoline()
func Getpgrp() (pgrp int) {
	r0, _, _ := syscall_rawSyscall(funcPC(libc_getpgrp_trampoline), 0, 0, 0)
	pgrp = int(r0)
	return
}
func libc_getpgrp_trampoline()
func Getpid() (pid int) {
	r0, _, _ := syscall_rawSyscall(funcPC(libc_getpid_trampoline), 0, 0, 0)
	pid = int(r0)
	return
}
func libc_getpid_trampoline()
func Getppid() (ppid int) {
	r0, _, _ := syscall_rawSyscall(funcPC(libc_getppid_trampoline), 0, 0, 0)
	ppid = int(r0)
	return
}
func libc_getppid_trampoline()
func Getpriority(which int, who int) (prio int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_getpriority_trampoline), uintptr(which), uintptr(who), 0)
	prio = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getpriority_trampoline()
func Getrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_getrlimit_trampoline), uintptr(which), uintptr(unsafe.Pointer(lim)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getrlimit_trampoline()
func Getrusage(who int, rusage *Rusage) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_getrusage_trampoline), uintptr(who), uintptr(unsafe.Pointer(rusage)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getrusage_trampoline()
func Getsid(pid int) (sid int, err error) {
	r0, _, e1 := syscall_rawSyscall(funcPC(libc_getsid_trampoline), uintptr(pid), 0, 0)
	sid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_getsid_trampoline()
func Getuid() (uid int) {
	r0, _, _ := syscall_rawSyscall(funcPC(libc_getuid_trampoline), 0, 0, 0)
	uid = int(r0)
	return
}
func libc_getuid_trampoline()
func Issetugid() (tainted bool) {
	r0, _, _ := syscall_rawSyscall(funcPC(libc_issetugid_trampoline), 0, 0, 0)
	tainted = bool(r0 != 0)
	return
}
func libc_issetugid_trampoline()
func Kqueue() (fd int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_kqueue_trampoline), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_kqueue_trampoline()
func Lchown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_lchown_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_lchown_trampoline()
func Link(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_link_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_link_trampoline()
func Linkat(pathfd int, path string, linkfd int, link string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall6(funcPC(libc_linkat_trampoline), uintptr(pathfd), uintptr(unsafe.Pointer(_p0)), uintptr(linkfd), uintptr(unsafe.Pointer(_p1)), uintptr(flags), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_linkat_trampoline()
func Listen(s int, backlog int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_listen_trampoline), uintptr(s), uintptr(backlog), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_listen_trampoline()
func Mkdir(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_mkdir_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_mkdir_trampoline()
func Mkdirat(dirfd int, path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_mkdirat_trampoline), uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_mkdirat_trampoline()
func Mkfifo(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_mkfifo_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_mkfifo_trampoline()
func Mknod(path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_mknod_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_mknod_trampoline()
func Open(path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := syscall_syscall(funcPC(libc_open_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm))
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_open_trampoline()
func Openat(dirfd int, path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := syscall_syscall6(funcPC(libc_openat_trampoline), uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm), 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_openat_trampoline()
func Pathconf(path string, name int) (val int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := syscall_syscall(funcPC(libc_pathconf_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(name), 0)
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_pathconf_trampoline()
func libc_pread_trampoline()
func libc_pwrite_trampoline()
func read(fd int, p []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := syscall_syscall(funcPC(libc_read_trampoline), uintptr(fd), uintptr(_p0), uintptr(len(p)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_read_trampoline()
func Readlink(path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 unsafe.Pointer
	if len(buf) > 0 {
		_p1 = unsafe.Pointer(&buf[0])
	} else {
		_p1 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := syscall_syscall(funcPC(libc_readlink_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(_p1), uintptr(len(buf)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_readlink_trampoline()
func Readlinkat(dirfd int, path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 unsafe.Pointer
	if len(buf) > 0 {
		_p1 = unsafe.Pointer(&buf[0])
	} else {
		_p1 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := syscall_syscall6(funcPC(libc_readlinkat_trampoline), uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(_p1), uintptr(len(buf)), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_readlinkat_trampoline()
func Rename(from string, to string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(from)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(to)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_rename_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_rename_trampoline()
func Renameat(fromfd int, from string, tofd int, to string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(from)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(to)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall6(funcPC(libc_renameat_trampoline), uintptr(fromfd), uintptr(unsafe.Pointer(_p0)), uintptr(tofd), uintptr(unsafe.Pointer(_p1)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_renameat_trampoline()
func Revoke(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_revoke_trampoline), uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_revoke_trampoline()
func Rmdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_rmdir_trampoline), uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_rmdir_trampoline()
func libc_lseek_trampoline()
func Select(n int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (err error) {
	_, _, e1 := syscall_syscall6(funcPC(libc_select_trampoline), uintptr(n), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_select_trampoline()
func Setegid(egid int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_setegid_trampoline), uintptr(egid), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setegid_trampoline()
func Seteuid(euid int) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_seteuid_trampoline), uintptr(euid), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_seteuid_trampoline()
func Setgid(gid int) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_setgid_trampoline), uintptr(gid), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setgid_trampoline()
func Setlogin(name string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(name)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_setlogin_trampoline), uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setlogin_trampoline()
func Setpgid(pid int, pgid int) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_setpgid_trampoline), uintptr(pid), uintptr(pgid), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setpgid_trampoline()
func Setpriority(which int, who int, prio int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_setpriority_trampoline), uintptr(which), uintptr(who), uintptr(prio))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setpriority_trampoline()
func Setprivexec(flag int) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_setprivexec_trampoline), uintptr(flag), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setprivexec_trampoline()
func Setregid(rgid int, egid int) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_setregid_trampoline), uintptr(rgid), uintptr(egid), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setregid_trampoline()
func Setreuid(ruid int, euid int) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_setreuid_trampoline), uintptr(ruid), uintptr(euid), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setreuid_trampoline()
func Setrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_setrlimit_trampoline), uintptr(which), uintptr(unsafe.Pointer(lim)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setrlimit_trampoline()
func Setsid() (pid int, err error) {
	r0, _, e1 := syscall_rawSyscall(funcPC(libc_setsid_trampoline), 0, 0, 0)
	pid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setsid_trampoline()
func Settimeofday(tp *Timeval) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_settimeofday_trampoline), uintptr(unsafe.Pointer(tp)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_settimeofday_trampoline()
func Setuid(uid int) (err error) {
	_, _, e1 := syscall_rawSyscall(funcPC(libc_setuid_trampoline), uintptr(uid), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_setuid_trampoline()
func Symlink(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_symlink_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_symlink_trampoline()
func Symlinkat(oldpath string, newdirfd int, newpath string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(oldpath)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(newpath)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_symlinkat_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(newdirfd), uintptr(unsafe.Pointer(_p1)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_symlinkat_trampoline()
func Sync() (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_sync_trampoline), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_sync_trampoline()
func libc_truncate_trampoline()
func Umask(newmask int) (oldmask int) {
	r0, _, _ := syscall_syscall(funcPC(libc_umask_trampoline), uintptr(newmask), 0, 0)
	oldmask = int(r0)
	return
}
func libc_umask_trampoline()
func Undelete(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_undelete_trampoline), uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_undelete_trampoline()
func Unlink(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_unlink_trampoline), uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_unlink_trampoline()
func Unlinkat(dirfd int, path string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_unlinkat_trampoline), uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(flags))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_unlinkat_trampoline()
func Unmount(path string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall(funcPC(libc_unmount_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(flags), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_unmount_trampoline()
func write(fd int, p []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := syscall_syscall(funcPC(libc_write_trampoline), uintptr(fd), uintptr(_p0), uintptr(len(p)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_write_trampoline()
func libc_mmap_trampoline()
func munmap(addr uintptr, length uintptr) (err error) {
	_, _, e1 := syscall_syscall(funcPC(libc_munmap_trampoline), uintptr(addr), uintptr(length), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_munmap_trampoline()
func readlen(fd int, buf *byte, nbuf int) (n int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_read_trampoline), uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func writelen(fd int, buf *byte, nbuf int) (n int, err error) {
	r0, _, e1 := syscall_syscall(funcPC(libc_write_trampoline), uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
func libc_gettimeofday_trampoline()
