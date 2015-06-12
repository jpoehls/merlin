package main

import (
	"github.com/jpoehls/go-conduit"
	"os"
)

func openConduit() *conduit.Conn {
	// TODO: find .merconfig and get the "phabricator.uri"
	// TODO: find the user/cert in .merrc that corresponds with the phabricator.uri we found

	conn, err := conduit.Dial(os.Getenv("MER_HOST"))
	if err != nil {
		fatalf("error connecting to host: %v", err)
	}
	err = conn.Connect(os.Getenv("MER_USER"), os.Getenv("MER_CERT"))
	if err != nil {
		fatalf("error connecting session: %v", err)
	}
	return conn
}
