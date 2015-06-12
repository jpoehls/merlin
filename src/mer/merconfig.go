package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type merconfig struct {
	PhabricatorURI string `json:"phabricator.uri"`
}

func merconfigPath() string {
	wd, err := os.Getwd()
	if err != nil {
		fatalf("error getting pwd: %s", err)
	}
	p := filepath.Join(wd, ".merconfig")
	return p
}

func openMerconfig() *merconfig {
	p := merconfigPath()
	f, err := ioutil.ReadFile(p)
	if err != nil && !os.IsNotExist(err) {
		fatalf("error reading %s: %s", p, err)
	}

	var o merconfig
	err = json.Unmarshal(f, &o)
	if err != nil {
		fatalf("error parsing %s: %s", p, err)
	}
	return &o
}
