package main

import (
	"os"
	"poehls.me/go-conduit"
)

func openConduit() *conduit.Conn {
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
