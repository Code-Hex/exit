package exit

// Auto generated by `mkerrs` in internal/cmd/mkerrs/main.go
// Do not edit this.
const (
	OK = 0

	USAGE = iota + 63
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

type base struct {
	err  error
	code int
}

// OK means the successful exit is always indicated by a status of 0.
type ok struct{ base }

func (o ok) ExitCode() int { return o.code }
func (o ok) Error() string { return o.err.Error() }

// MakeOK will return ok struct as error
func MakeOK(err error) error { return ok{base{err: err, code: OK}} }

// USAGE means the command was used incorrectly, e.g., with the
// wrong number of arguments, a bad flag, a bad syntax
// in a parameter, or whatever
type usage struct{ base }

func (u usage) ExitCode() int { return u.code }
func (u usage) Error() string { return u.err.Error() }

// MakeUsage will return usage struct as error
func MakeUsage(err error) error { return usage{base{err: err, code: USAGE}} }

// DATAERR means the input data was incorrect in some way. This
// should only be used for user's data and not system
// files
type dataerr struct{ base }

func (d dataerr) ExitCode() int { return d.code }
func (d dataerr) Error() string { return d.err.Error() }

// MakeDataErr will return dataerr struct as error
func MakeDataErr(err error) error { return dataerr{base{err: err, code: DATAERR}} }

// NOINPUT means an input file (not a system file) did not exist or
// was not readable. This could also include errors
// like `No message` to a mailer (if it cared to
// catch it).
type noinput struct{ base }

func (n noinput) ExitCode() int { return n.code }
func (n noinput) Error() string { return n.err.Error() }

// MakeNoInput will return noinput struct as error
func MakeNoInput(err error) error { return noinput{base{err: err, code: NOINPUT}} }

// NOUSER means the user specified did not exist. This might be
// used for mail addresses or remote logins.
type nouser struct{ base }

func (n nouser) ExitCode() int { return n.code }
func (n nouser) Error() string { return n.err.Error() }

// MakeNoUser will return nouser struct as error
func MakeNoUser(err error) error { return nouser{base{err: err, code: NOUSER}} }

// NOHOST means the host specified did not exist. This is used in
// mail addresses or network requests.
type nohost struct{ base }

func (n nohost) ExitCode() int { return n.code }
func (n nohost) Error() string { return n.err.Error() }

// MakeNoHost will return nohost struct as error
func MakeNoHost(err error) error { return nohost{base{err: err, code: NOHOST}} }

// UNAVAILABLE means a service is unavailable. This can occur if a support
// program or file does not exist. This can also
// be used as a catchall message when something you
// wanted to do doesn't work, but you don't know why.
type unavailable struct{ base }

func (u unavailable) ExitCode() int { return u.code }
func (u unavailable) Error() string { return u.err.Error() }

// MakeUnAvailable will return unavailable struct as error
func MakeUnAvailable(err error) error { return unavailable{base{err: err, code: UNAVAILABLE}} }

// SOFTWARE means an internal software error has been detected. This
// should be limited to non-operating system related
// errors as possible.
type software struct{ base }

func (s software) ExitCode() int { return s.code }
func (s software) Error() string { return s.err.Error() }

// MakeSoftWare will return software struct as error
func MakeSoftWare(err error) error { return software{base{err: err, code: SOFTWARE}} }

// OSERR means an operating system error has been detected. This
// is intended to be used for such things as `cannot
// fork`, `cannot create pipe`, or the like. It
// includes things like getuid returning a user that
// does not exist in the passwd file.
type oserr struct{ base }

func (o oserr) ExitCode() int { return o.code }
func (o oserr) Error() string { return o.err.Error() }

// MakeOSErr will return oserr struct as error
func MakeOSErr(err error) error { return oserr{base{err: err, code: OSERR}} }

// OSFILE means some system file (e.g., /etc/passwd, /var/run/utmp,
// etc.) does not exist, cannot be opened, or has some
// sort of error (e.g., syntax error).
type osfile struct{ base }

func (o osfile) ExitCode() int { return o.code }
func (o osfile) Error() string { return o.err.Error() }

// MakeOSFile will return osfile struct as error
func MakeOSFile(err error) error { return osfile{base{err: err, code: OSFILE}} }

// CANTCREAT means a (user specified) output file cannot be created.
type cantcreat struct{ base }

func (c cantcreat) ExitCode() int { return c.code }
func (c cantcreat) Error() string { return c.err.Error() }

// MakeCantCreate will return cantcreat struct as error
func MakeCantCreate(err error) error { return cantcreat{base{err: err, code: CANTCREAT}} }

// IOERR means an error occurred while doing I/O on some file.
type ioerr struct{ base }

func (i ioerr) ExitCode() int { return i.code }
func (i ioerr) Error() string { return i.err.Error() }

// MakeIOErr will return ioerr struct as error
func MakeIOErr(err error) error { return ioerr{base{err: err, code: IOERR}} }

// TEMPFAIL means temporary failure, indicating something that is not
// really an error. In sendmail, this means that a
// mailer (e.g.) could not create a connection, and
// the request should be reattempted later.
type tempfail struct{ base }

func (t tempfail) ExitCode() int { return t.code }
func (t tempfail) Error() string { return t.err.Error() }

// MakeTempFail will return tempfail struct as error
func MakeTempFail(err error) error { return tempfail{base{err: err, code: TEMPFAIL}} }

// PROTOCOL means the remote system returned something that was `not
// possible` during a protocol exchange.
type protocol struct{ base }

func (p protocol) ExitCode() int { return p.code }
func (p protocol) Error() string { return p.err.Error() }

// MakeProtocol will return protocol struct as error
func MakeProtocol(err error) error { return protocol{base{err: err, code: PROTOCOL}} }

// NOPERM means you did not have sufficient permission to perform
// the operation.  This is not intended for file sys­tem
// problems, which should use NOINPUT or
// CANTCREAT, but rather for higher level permis­sions.
type noperm struct{ base }

func (n noperm) ExitCode() int { return n.code }
func (n noperm) Error() string { return n.err.Error() }

// MakeNoPerm will return noperm struct as error
func MakeNoPerm(err error) error { return noperm{base{err: err, code: NOPERM}} }

// CONFIG means Something was found in an unconfigured or
// miscon­figured state.
type config struct{ base }

func (c config) ExitCode() int { return c.code }
func (c config) Error() string { return c.err.Error() }

// MakeConfig will return config struct as error
func MakeConfig(err error) error { return config{base{err: err, code: CONFIG}} }
