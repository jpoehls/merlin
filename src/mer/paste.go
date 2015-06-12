package main

import (
	"github.com/jpoehls/go-conduit"
	"io/ioutil"
	"os"
	"strconv"
)

var cmdPaste = &Command{
	UsageLines: []string{
		"paste id",
		"paste [-title title] [-lang language]",
	},
	Short: "retrieves or shares text using the Paste application",
	Long: `
Share and grab text using the Paste application.
To create a paste, use stdin to provide the text:

  $ cat list_of_ducks.txt | mer paste

To retrieve a paste, specify the paste ID:

  $ mer paste P123

-lang language
    Language for syntax highlighting.

-title title
    Title for the paste.
`,
}

func init() {
	cmdPaste.Run = runPaste // break init cycle
}

var pasteJSON = cmdPaste.Flag.Bool("json", false, "")
var pasteTitle = cmdPaste.Flag.String("title", "", "")
var pasteLang = cmdPaste.Flag.String("lang", "", "")

func runPaste(cmd *Command, args []string) {
	if len(args) > 0 {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fatalf("id must be a number: %v", err)
		}

		if id != 0 {
			paste := getPaste(id)
			writePaste(paste)
		}

		return
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fatalf("error reading stdin: %v", err)
	}

	paste := createPaste(bytes)
	writePaste(paste)
}

func getPaste(id uint64) *conduit.PasteItem {
	conn := openConduit()
	qresp, err := conn.PasteQuery(&conduit.PasteQueryParams{
		IDs: []uint64{1},
	})
	if err != nil {
		fatalf("error getting pastes: %v", err)
	}

	if len(qresp) == 0 {
		fatalf("paste not found")
	}

	return qresp[0]
}

func createPaste(bytes []byte) *conduit.PasteItem {
	conn := openConduit()
	params := &conduit.PasteCreateParams{
		Content: string(bytes),
	}
	if *pasteTitle != "" {
		params.Title = *pasteTitle
	}
	if *pasteLang != "" {
		params.Language = *pasteLang
	}
	paste, err := conn.PasteCreate(params)
	if err != nil {
		fatalf("error creating paste: %v", err)
	}

	return paste
}

func writePaste(p *conduit.PasteItem) {
	if _, err := os.Stdout.WriteString("Pasted as " + p.ObjectName + "\n" + p.URI + "\n"); err != nil {
		fatalf("error writing paste to stdout: %v", err)
	}
}
