package main

import (
	"fmt"
)

var cmdInstallCert = &Command{
	Run:        runInstallCert,
	UsageLines: []string{"install-certificate [uri]"},
	Short:      "installs conduit credentials",
	Long: `
Installs conduit credentials into your ~/.merrc for the given
install of Phabricator. You need to do this before you can use 'mer',
as it enables 'mer' to link your command-line activity with
your account on the web.

Run this command from within a project directory to install that
project's certificate, or specify an explicit URI
(like "https://phabricator.example.com").
`,
}

func runInstallCert(cmd *Command, args []string) {
	// TODO: implement what we claim to do

	fmt.Printf("URI: %v\n", args)
}
