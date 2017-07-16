package main

import (
	"os"
	"testing"
)

/*
 * The unit tests in this file simulate command line invocation.
 */

func TestMain(test *testing.T) {
	os.Args = []string{programName}
	main()
}

func TestMainWithArguments(test *testing.T) {
	os.Args = []string{
		programName,
		"--flagKey",
		"From command option",
		"--flagKeyOnly",
		"From command option",
		"--setKey",
		"From command option",
	}
	main()
}
