package main

import (
	"fmt"
)

var cmdPaste = &Command{
	UsageLines: []string{
		"paste id [--json]",
		"paste [--title title] [--lang language] [--json]",
	},
	Short: "retrieves or shares text using the Paste application",
	Long: `
Share and grab text using the Paste application.
To create a paste, use stdin to provide the text:

  $ cat list_of_ducks.txt | mer paste

To retrieve a paste, specify the paste ID:

  $ mer paste P123

--json
    Output in JSON format.

--lang language
    Language for syntax highlighting.

--title title
    Title for the paste.
`,
}

func init() {
	cmdPaste.Run = runPaste // break init cycle
}

var pasteID = cmdPaste.Flag.Uint64("id", 0, "")
var pasteJSON = cmdPaste.Flag.Bool("json", false, "")
var pasteTitle = cmdPaste.Flag.String("title", "", "")
var pasteLang = cmdPaste.Flag.String("lang", "", "")

func runPaste(cmd *Command, args []string) {
	// TODO: implement what we claim to do

	fmt.Printf("ID: %d\n", *pasteID)
}
