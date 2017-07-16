package main

import (
	"fmt"

	"github.com/docktermj/go-hello-viper/configuration"
	"github.com/docktermj/go-hello-viper/etc"
	"github.com/docopt/docopt-go"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

func main() {

	usage := `
Usage:
    go-hello-viper [options]

Options:
   -h, --help            Show this help
   --flagKey=<text>      Sample text.
   --flagKeyOnly=<text>  Sample text.
   --setKey=<text>       Sample text.
   --debug               Additional debugging statements.
   --version             Show version information
`

	// DocOpt processing.

	commandVersion := fmt.Sprintf("%s %s-%s", programName, buildVersion, buildIteration)
	args, _ := docopt.Parse(usage, nil, true, commandVersion, false)

	configuration.LoadConfig(args)
	etc.TrySomething()
}
