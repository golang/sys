// Generated code. DO NOT EDIT.

package unix

import (
	"syscall"
	"unsafe"
)

func pipe(p *[2]_C_int) (n int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procpipe)), 1, uintptr(unsafe.Pointer(p)), 0, 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procgetsockname)), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Getcwd(buf []byte) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procGetcwd)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func getgroups(ngid int, gid *_Gid_t) (n int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procgetgroups)), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func setgroups(ngid int, gid *_Gid_t) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procsetgroups)), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func wait4(pid int32, statusp *_C_int, options int, rusage *Rusage) (wpid int32, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procwait4)), 4, uintptr(pid), uintptr(unsafe.Pointer(statusp)), uintptr(options), uintptr(unsafe.Pointer(rusage)), 0, 0)
	wpid = int32(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func gethostname(buf []byte) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procgethostname)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func utimes(path string, times *[2]Timeval) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procutimes)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func utimensat(fd int, path string, times *[2]Timespec, flag int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procutimensat)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), uintptr(flag), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procfcntl)), 3, uintptr(fd), uintptr(cmd), uintptr(arg), 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func futimesat(fildes int, path *byte, times *[2]Timeval) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procfutimesat)), 3, uintptr(fildes), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(times)), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procaccept)), 3, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_recvmsg)), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_sendmsg)), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func acct(path *byte) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procacct)), 1, uintptr(unsafe.Pointer(path)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func __makedev(version int, major uint, minor uint) (val uint64) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&proc__makedev)), 3, uintptr(version), uintptr(major), uintptr(minor), 0, 0, 0)
	val = uint64(r0)
	return
}
func __major(version int, dev uint64) (val uint) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&proc__major)), 2, uintptr(version), uintptr(dev), 0, 0, 0, 0)
	val = uint(r0)
	return
}
func __minor(version int, dev uint64) (val uint) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&proc__minor)), 2, uintptr(version), uintptr(dev), 0, 0, 0, 0)
	val = uint(r0)
	return
}
func ioctl(fd int, req uint, arg uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procioctl)), 3, uintptr(fd), uintptr(req), uintptr(arg), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func poll(fds *PollFd, nfds int, timeout int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procpoll)), 3, uintptr(unsafe.Pointer(fds)), uintptr(nfds), uintptr(timeout), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Access(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procAccess)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Adjtime(delta *Timeval, olddelta *Timeval) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procAdjtime)), 2, uintptr(unsafe.Pointer(delta)), uintptr(unsafe.Pointer(olddelta)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Chdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChdir)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Chmod(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChmod)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Chown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChown)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Chroot(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChroot)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Close(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procClose)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Creat(path string, mode uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procCreat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Dup(fd int) (nfd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procDup)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	nfd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Dup2(oldfd int, newfd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procDup2)), 2, uintptr(oldfd), uintptr(newfd), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Exit(code int) {
	sysvicall6(uintptr(unsafe.Pointer(&procExit)), 1, uintptr(code), 0, 0, 0, 0, 0)
	return
}
func Faccessat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFaccessat)), 4, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(flags), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fchdir(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchdir)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fchmod(fd int, mode uint32) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchmod)), 2, uintptr(fd), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchmodat)), 4, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(flags), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fchown(fd int, uid int, gid int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchown)), 3, uintptr(fd), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchownat)), 5, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), uintptr(flags), 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fdatasync(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFdatasync)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Flock(fd int, how int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFlock)), 2, uintptr(fd), uintptr(how), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fpathconf(fd int, name int) (val int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFpathconf)), 2, uintptr(fd), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fstat(fd int, stat *Stat_t) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFstat)), 2, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fstatat(fd int, path string, stat *Stat_t, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFstatat)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), uintptr(flags), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fstatvfs(fd int, vfsstat *Statvfs_t) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFstatvfs)), 2, uintptr(fd), uintptr(unsafe.Pointer(vfsstat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Getdents(fd int, buf []byte, basep *uintptr) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procGetdents)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(unsafe.Pointer(basep)), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Getgid() (gid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&procGetgid)), 0, 0, 0, 0, 0, 0, 0)
	gid = int(r0)
	return
}
func Getpid() (pid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&procGetpid)), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	return
}
func Getpgid(pid int) (pgid int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGetpgid)), 1, uintptr(pid), 0, 0, 0, 0, 0)
	pgid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Getpgrp() (pgid int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGetpgrp)), 0, 0, 0, 0, 0, 0, 0)
	pgid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Geteuid() (euid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procGeteuid)), 0, 0, 0, 0, 0, 0, 0)
	euid = int(r0)
	return
}
func Getegid() (egid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procGetegid)), 0, 0, 0, 0, 0, 0, 0)
	egid = int(r0)
	return
}
func Getppid() (ppid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procGetppid)), 0, 0, 0, 0, 0, 0, 0)
	ppid = int(r0)
	return
}
func Getpriority(which int, who int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procGetpriority)), 2, uintptr(which), uintptr(who), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Getrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGetrlimit)), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Getrusage(who int, rusage *Rusage) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGetrusage)), 2, uintptr(who), uintptr(unsafe.Pointer(rusage)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Gettimeofday(tv *Timeval) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGettimeofday)), 1, uintptr(unsafe.Pointer(tv)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Getuid() (uid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&procGetuid)), 0, 0, 0, 0, 0, 0, 0)
	uid = int(r0)
	return
}
func Kill(pid int, signum syscall.Signal) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procKill)), 2, uintptr(pid), uintptr(signum), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Lchown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procLchown)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
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
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procLink)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Listen(s int, backlog int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_llisten)), 2, uintptr(s), uintptr(backlog), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Lstat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procLstat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Madvise(b []byte, advice int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMadvise)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), uintptr(advice), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Mkdir(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMkdir)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Mkdirat(dirfd int, path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMkdirat)), 3, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Mkfifo(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMkfifo)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Mkfifoat(dirfd int, path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMkfifoat)), 3, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Mknod(path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMknod)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Mknodat(dirfd int, path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMknodat)), 4, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Mlock(b []byte) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMlock)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Mlockall(flags int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMlockall)), 1, uintptr(flags), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Mprotect(b []byte, prot int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMprotect)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), uintptr(prot), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Msync(b []byte, flags int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMsync)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), uintptr(flags), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Munlock(b []byte) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMunlock)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Munlockall() (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMunlockall)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Nanosleep(time *Timespec, leftover *Timespec) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procNanosleep)), 2, uintptr(unsafe.Pointer(time)), uintptr(unsafe.Pointer(leftover)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Open(path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procOpen)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Openat(dirfd int, path string, flags int, mode uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procOpenat)), 4, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(flags), uintptr(mode), 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Pathconf(path string, name int) (val int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPathconf)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Pause() (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPause)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Pread(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPread)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPwrite)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func read(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procread)), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Readlink(path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	if len(buf) > 0 {
		_p1 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procReadlink)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(len(buf)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
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
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procRename)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error) {
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
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procRenameat)), 4, uintptr(olddirfd), uintptr(unsafe.Pointer(_p0)), uintptr(newdirfd), uintptr(unsafe.Pointer(_p1)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Rmdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procRmdir)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proclseek)), 3, uintptr(fd), uintptr(offset), uintptr(whence), 0, 0, 0)
	newoffset = int64(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Select(n int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSelect)), 5, uintptr(n), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Setegid(egid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetegid)), 1, uintptr(egid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Seteuid(euid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSeteuid)), 1, uintptr(euid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Setgid(gid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetgid)), 1, uintptr(gid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Sethostname(p []byte) (err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSethostname)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Setpgid(pid int, pgid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetpgid)), 2, uintptr(pid), uintptr(pgid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Setpriority(which int, who int, prio int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSetpriority)), 3, uintptr(which), uintptr(who), uintptr(prio), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Setregid(rgid int, egid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetregid)), 2, uintptr(rgid), uintptr(egid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Setreuid(ruid int, euid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetreuid)), 2, uintptr(ruid), uintptr(euid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Setrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetrlimit)), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Setsid() (pid int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetsid)), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Setuid(uid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetuid)), 1, uintptr(uid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Shutdown(s int, how int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procshutdown)), 2, uintptr(s), uintptr(how), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Stat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procStat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Statvfs(path string, vfsstat *Statvfs_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procStatvfs)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(vfsstat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
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
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSymlink)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Sync() (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSync)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Times(tms *Tms) (ticks uintptr, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procTimes)), 1, uintptr(unsafe.Pointer(tms)), 0, 0, 0, 0, 0)
	ticks = uintptr(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Truncate(path string, length int64) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procTruncate)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Fsync(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFsync)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Ftruncate(fd int, length int64) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFtruncate)), 2, uintptr(fd), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Umask(mask int) (oldmask int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procUmask)), 1, uintptr(mask), 0, 0, 0, 0, 0)
	oldmask = int(r0)
	return
}
func Uname(buf *Utsname) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procUname)), 1, uintptr(unsafe.Pointer(buf)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Unmount(target string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(target)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procumount)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(flags), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Unlink(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procUnlink)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Unlinkat(dirfd int, path string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procUnlinkat)), 3, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(flags), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Ustat(dev int, ubuf *Ustat_t) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procUstat)), 2, uintptr(dev), uintptr(unsafe.Pointer(ubuf)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func Utime(path string, buf *Utimbuf) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procUtime)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(buf)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_bind)), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_connect)), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func mmap(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procmmap)), 6, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flag), uintptr(fd), uintptr(pos))
	ret = uintptr(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func munmap(addr uintptr, length uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procmunmap)), 2, uintptr(addr), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procsendfile)), 4, uintptr(outfd), uintptr(infd), uintptr(unsafe.Pointer(offset)), uintptr(count), 0, 0)
	written = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_sendto)), 6, uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		err = e1
	}
	return
}
func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_socket)), 3, uintptr(domain), uintptr(typ), uintptr(proto), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&proc__xnet_socketpair)), 4, uintptr(domain), uintptr(typ), uintptr(proto), uintptr(unsafe.Pointer(fd)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func write(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procwrite)), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_getsockopt)), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procgetpeername)), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procsetsockopt)), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0)
	if e1 != 0 {
		err = e1
	}
	return
}
func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procrecvfrom)), 6, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
