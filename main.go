package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"text/template"
	"unicode"
	"unicode/utf8"
)

// A Command is an implementation of a mer command
// like mer call-conduit or mer paste.
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)

	// UsageLines is the collection of one-line usage messages.
	// The first word in the first line is taken to be the command name.
	UsageLines []string

	// Short is the short description shown in the 'mer help' output.
	Short string

	// Long is the long message shown in the 'mer help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own
	// flag parsing.
	CustomFlags bool
}

// Name returns the command's name: the first word in the first usage line.
func (c *Command) Name() string {
	name := c.UsageLines[0]
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// Usage writes the command's usage information to stderr and exits.
func (c *Command) Usage() {

	for i, l := range c.UsageLines {
		if i == 0 {
			fmt.Fprintf(os.Stderr, "usage: %s\n\n", l)
		} else {
			fmt.Fprintf(os.Stderr, "       %s\n\n", l)
		}
	}

	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}

// Runnable reports whether the command can be run; otherwise
// it is a documentation pseudo-command such as importpath.
func (c *Command) Runnable() bool {
	return c.Run != nil
}

// Commands lists the available commands and help topics.
// The order here is the order in which they are printed by 'mer help'.
var commands = []*Command{
	cmdCallConduit,
	cmdPaste,

	helpCert,
}

var exitStatus = 0
var exitMu sync.Mutex

func setExitStatus(n int) {
	exitMu.Lock()
	if exitStatus < n {
		exitStatus = n
	}
	exitMu.Unlock()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			exit()
			return
		}
	}

	fmt.Fprintf(os.Stderr, "mer: unknown subcommand %q\nRun 'mer help' for usage.\n", args[0])
	setExitStatus(2)
	exit()
}

var usageTemplate = `Merlin is a tool for working with Phabricator.

Usage:

	mer command [arguments]

The commands are:
{{range .}}{{if .Runnable}}
    {{.Name | printf "%-13s"}} {{.Short}}{{end}}{{end}}

Use "mer help [command]" for more information about a command.

Additional help topics:
{{range .}}{{if not .Runnable}}
    {{.Name | printf "%-13s"}} {{.Short}}{{end}}{{end}}

Use "mer help [topic]" for more information about that topic.

`

var helpTemplate = `{{if .Runnable}}usage: {{range $i, $l := .UsageLines}}{{if gt $i 0}}       {{end}}mer {{$l}}
{{end}}
{{end}}{{.Long | trim}}
`

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{
		"trim":       strings.TrimSpace,
		"capitalize": capitalize})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}

func printUsage(w io.Writer) {
	tmpl(w, usageTemplate, commands)
}

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

// help implements the 'help' command.
func help(args []string) {
	if len(args) == 0 {
		printUsage(os.Stdout)
		// not exit 2: succeeded at 'mer help'.
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: mer help command\n\nToo many arguments given.\n")
		os.Exit(2) // failed at 'mer help'
	}

	arg := args[0]

	for _, cmd := range commands {
		if cmd.Name() == arg {
			tmpl(os.Stdout, helpTemplate, cmd)
			// not exit 2: succeeded at 'mer help cmd'.
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic %#q.  Run 'mer help'.\n", arg)
	os.Exit(2) // failed at 'mer help cmd'
}

func exit() {
	os.Exit(exitStatus)
}

func fatalf(format string, args ...interface{}) {
	errorf(format, args...)
	exit()
}

func errorf(format string, args ...interface{}) {
	logf(format, args...)
	setExitStatus(1)
}

var logf = log.Printf

func exitIfErrors() {
	if exitStatus != 0 {
		exit()
	}
}
