package main

import (
	"fmt"
	"runtime"
)

var cmdVersion = &Command{
	Run:        runVersion,
	UsageLines: []string{"version"},
	Short:      "print Merlin version",
	Long:       `Version prints the Merlin version.`,
}

func runVersion(cmd *Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}

	// TODO: Wire up a real version number.
	fmt.Printf("Merlin version %s\n%s %s/%s\n", "0.0.0", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
