# exit
Error makes exit code meaningful

[![GoDoc](https://godoc.org/github.com/Code-Hex/exit?status.svg)](https://godoc.org/github.com/Code-Hex/exit) 
[![Build Status](https://travis-ci.org/Code-Hex/exit.svg?branch=master)](https://travis-ci.org/Code-Hex/exit) 
[![Coverage Status](https://coveralls.io/repos/github/Code-Hex/exit/badge.svg?branch=master)](https://coveralls.io/github/Code-Hex/exit?branch=master) 
[![Go Report Card](https://goreportcard.com/badge/github.com/Code-Hex/exit)](https://goreportcard.com/report/github.com/Code-Hex/exit)

# Description
Inspired from `sysexits.h` and created it.  
The following are quoted contents from manpage of the `sysexits.h`.
> According to style(9), it is not a good practice to call exit(3) with
arbitrary values to indicate a failure condition when ending a program.
Instead, the pre-defined exit codes from sysexits should be used, so the
caller of the process can get a rough estimation about the failure class
without looking up the source code.  
  
You can read [example](https://github.com/Code-Hex/exit/blob/master/eg/main.go).  
And if you want to know more details, please read [here](https://github.com/Code-Hex/exit/blob/master/errors.go).

# Install
    go get github.com/Code-Hex/exit

# Contributing
Welcome!!  
I'm waiting for pull requests or reporting issues.

# License
[MIT](https://github.com/Code-Hex/sigctx/blob/master/LICENSE)

# Author
[CodeHex](https://twitter.com/CodeHex)  
