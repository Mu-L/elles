//go:build openbsd

package os2

import "golang.org/x/sys/unix"

func statfs(m string) (int, error) {
	var vfs unix.Statfs_t
	err := unix.Statfs(m, &vfs)
	return int(vfs.F_bsize), err
}
