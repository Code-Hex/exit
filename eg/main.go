package main

import (
	"fmt"
	"os"

	"github.com/Code-Hex/exit"
	"github.com/pkg/errors"
)

type exiter interface {
	ExitCode() int
}

type causer interface {
	Cause() error
}

func unwrapErrors(err error) (int, error) {
	for e := err; e != nil; {
		switch e.(type) {
		case exiter:
			return e.(exiter).ExitCode(), e
		case causer:
			e = e.(causer).Cause()
		default:
			return 1, e // default error
		}
	}
	return 0, nil
}

func getExitCode(err error) int {
	for e := err; e != nil; {
		switch e.(type) {
		case exiter:
			return e.(exiter).ExitCode()
		case causer:
			e = e.(causer).Cause()
		}
	}
	return exit.OK
}

func fileoperate() error {
	if err := dosomething1(); err != nil {
		return err
	}
	return nil
}

func dosomething3() error {
	return exit.MakeOSFILE("Failed to operate files")
}

func dosomething2() error {
	if err := dosomething3(); err != nil {
		return errors.Wrap(err, "Second error")
	}
	return nil
}

func dosomething1() error {
	if err := dosomething2(); err != nil {
		return errors.Wrap(err, "Third error")
	}
	return nil
}

func run(trace bool) int {
	if err := fileoperate(); err != nil {
		var code int
		if trace {
			code = getExitCode(err)
			fmt.Printf("%+v\n", err)
		} else {
			code, err = unwrapErrors(err)
			fmt.Printf("%s\n", err.Error())
		}
		return code
	}

	return exit.OK
}

func main() {
	os.Exit(run(true))
}
