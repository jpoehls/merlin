# Merlin

A command line interface for [Phabricator](http://phabricator.org) written in Go. Aspires to be a viable replacement for [Arcanist](http://www.phabricator.com/docs/arcanist/) with all the benefits of being deployed as a static executable. I.e., no PHP!

Merlin (or `mer`, the name of the executable) relies heavily on [go-conduit](https://github.com/jpoehls/go-conduit), a Go package that handles communicating with Phabricator's [Conduit](https://secure.phabricator.com/book/phabdev/article/conduit/) API.

[Documentation](http://godoc.org/github.com/jpoehls/go-conduit)

# Installation

Assuming you have a working Go environment and `GOPATH/bin` is in your `PATH`, `mer` is a breeze to install:

`go get poehls.me/merlin`

Then verify that `mer` was installed correctly:

`mer -h`