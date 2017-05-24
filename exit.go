package exit

// http://kaworu.jpn.org/doc/FreeBSD/jman/man3/sysexits.3.php
const (
	OK = 0

	USAGE = iota + 64
	DATAERR
	NOINPUT
	NOUSER
	NOHOST
	UNAVAILABLE
	SOFTWARE
	OSERR
	OSFILE
	CANTCREAT
	IOERR
	TEMPFAIL
	PROTOCOL
	NOPERM
	CONFIG
)
