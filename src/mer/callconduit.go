package main

var cmdCallConduit = &Command{
	Run:        runCallConduit,
	UsageLines: []string{"call-conduit method"},
	Short:      "run a raw conduit method",
	Long: `
Allows you to make a raw Conduit method call:

  - Run this command from a working directory.
  - Call parameters are REQUIRED and read as a JSON blob from stdin.
  - Results are written to stdout as a JSON blob.

This workflow is primarily useful for writing scripts which
integrate with Phabricator. Examples:

  $ echo '{}' | mer call-conduit conduit.ping
  $ echo '{"phid":"PHID-FILE-xxxx"}' | mer call-conduit file.download
`,
}

func runCallConduit(cmd *Command, args []string) {
	// TODO: implement what we claim to do
}
