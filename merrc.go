package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type merrc struct {
	Hosts map[string]*merrcHost `json:"hosts"`
}

type merrcHost struct {
	User string `json:"user"`
	Cert string `json:"cert"`
}

func merrcPath() string {
	home := os.Getenv("HOME")
	p := filepath.Join(home, ".merrc")
	return p
}

func openMerrc() *merrc {
	p := merrcPath()
	f, err := ioutil.ReadFile(p)
	if err != nil && !os.IsNotExist(err) {
		errorf("error reading %s: %s", p, err)
	}

	var o merrc
	err = json.Unmarshal(f, &o)
	if err != nil {
		errorf("error parsing %s: %s", p, err)
	}
	return &o
}
