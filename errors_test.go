package exit

// Auto generated by `mkerrs` in internal/cmd/mkerrs/main.go
// Do not edit this.
import "testing"
import "errors"

func TestMakeOK(t *testing.T) {
	msg := "this is the error msg"
	err := MakeOK(errors.New(msg))
	if err.(ok).ExitCode() != OK {
		t.Fatal("Failed to get exitcode in TestMakeOK")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeUsage(t *testing.T) {
	msg := "this is the error msg"
	err := MakeUsage(errors.New(msg))
	if err.(usage).ExitCode() != USAGE {
		t.Fatal("Failed to get exitcode in TestMakeUSAGE")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeDataErr(t *testing.T) {
	msg := "this is the error msg"
	err := MakeDataErr(errors.New(msg))
	if err.(dataerr).ExitCode() != DATAERR {
		t.Fatal("Failed to get exitcode in TestMakeDATAERR")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeNoInput(t *testing.T) {
	msg := "this is the error msg"
	err := MakeNoInput(errors.New(msg))
	if err.(noinput).ExitCode() != NOINPUT {
		t.Fatal("Failed to get exitcode in TestMakeNOINPUT")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeNoUser(t *testing.T) {
	msg := "this is the error msg"
	err := MakeNoUser(errors.New(msg))
	if err.(nouser).ExitCode() != NOUSER {
		t.Fatal("Failed to get exitcode in TestMakeNOUSER")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeNoHost(t *testing.T) {
	msg := "this is the error msg"
	err := MakeNoHost(errors.New(msg))
	if err.(nohost).ExitCode() != NOHOST {
		t.Fatal("Failed to get exitcode in TestMakeNOHOST")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeUnAvailable(t *testing.T) {
	msg := "this is the error msg"
	err := MakeUnAvailable(errors.New(msg))
	if err.(unavailable).ExitCode() != UNAVAILABLE {
		t.Fatal("Failed to get exitcode in TestMakeUNAVAILABLE")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeSoftWare(t *testing.T) {
	msg := "this is the error msg"
	err := MakeSoftWare(errors.New(msg))
	if err.(software).ExitCode() != SOFTWARE {
		t.Fatal("Failed to get exitcode in TestMakeSOFTWARE")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeOSErr(t *testing.T) {
	msg := "this is the error msg"
	err := MakeOSErr(errors.New(msg))
	if err.(oserr).ExitCode() != OSERR {
		t.Fatal("Failed to get exitcode in TestMakeOSERR")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeOSFile(t *testing.T) {
	msg := "this is the error msg"
	err := MakeOSFile(errors.New(msg))
	if err.(osfile).ExitCode() != OSFILE {
		t.Fatal("Failed to get exitcode in TestMakeOSFILE")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeCantCreate(t *testing.T) {
	msg := "this is the error msg"
	err := MakeCantCreate(errors.New(msg))
	if err.(cantcreat).ExitCode() != CANTCREAT {
		t.Fatal("Failed to get exitcode in TestMakeCANTCREAT")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeIOErr(t *testing.T) {
	msg := "this is the error msg"
	err := MakeIOErr(errors.New(msg))
	if err.(ioerr).ExitCode() != IOERR {
		t.Fatal("Failed to get exitcode in TestMakeIOERR")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeTempFail(t *testing.T) {
	msg := "this is the error msg"
	err := MakeTempFail(errors.New(msg))
	if err.(tempfail).ExitCode() != TEMPFAIL {
		t.Fatal("Failed to get exitcode in TestMakeTEMPFAIL")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeProtocol(t *testing.T) {
	msg := "this is the error msg"
	err := MakeProtocol(errors.New(msg))
	if err.(protocol).ExitCode() != PROTOCOL {
		t.Fatal("Failed to get exitcode in TestMakePROTOCOL")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeNoPerm(t *testing.T) {
	msg := "this is the error msg"
	err := MakeNoPerm(errors.New(msg))
	if err.(noperm).ExitCode() != NOPERM {
		t.Fatal("Failed to get exitcode in TestMakeNOPERM")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}

func TestMakeConfig(t *testing.T) {
	msg := "this is the error msg"
	err := MakeConfig(errors.New(msg))
	if err.(config).ExitCode() != CONFIG {
		t.Fatal("Failed to get exitcode in TestMakeCONFIG")
	}
	if err.Error() != msg {
		t.Fatal("Unpexted error message, expected %s", msg)
	}
}
