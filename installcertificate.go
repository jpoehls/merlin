package main

import (
	"fmt"
)

var cmdInstallCert = &Command{
	UsageLines: []string{"install-certificate [uri]"},
	Short:      "installs conduit credentials",
	Long: `
Installs conduit credentials into your ~/.merrc for the given
install of Phabricator. You need to do this before you can use 'mer',
as it enables 'mer' to link your command-line activity with
your accunt on the web.

Run this command from within a project directory to install that
project's certificate, or specify an explicit URI
(like "https://phabricator.example.com").
`,
}

func init() {
	cmdInstallCert.Run = runInstallCert // break init cycle
}

var installCertURI = cmdPaste.Flag.String("uri", "", "")

func runInstallCert(cmd *Command, args []string) {
	// TODO: implement what we claim to do

	fmt.Printf("URI: %s\n", *installCertURI)
}
