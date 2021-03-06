# Merlin

A command line interface for [Phabricator](http://phabricator.org) written in Go. Aspires to be a viable replacement for [Arcanist](http://www.phabricator.com/docs/arcanist/) with all the benefits of being deployed as a static executable. I.e., no PHP!

Merlin (or `mer`, the name of the executable) relies heavily on [go-conduit](https://github.com/jpoehls/go-conduit), a Go package that handles communicating with Phabricator's [Conduit](https://secure.phabricator.com/book/phabdev/article/conduit/) API.

[Documentation](http://godoc.org/github.com/jpoehls/go-conduit)

# Building

Install [gb](http://getgb.io).

	go get github.com/constabulary/gb/...

Then build with `gb build` and grab the `bin\mer` executable.