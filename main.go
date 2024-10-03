package main

import (
	"hurricane/internal/master"
	"hurricane/internal/master/options"
)

type Info struct {
	GitCommit  string   `json:"git_commit"`
	GitVersion int      `json:"git_version"`
	GitArr     []string `json:"git_"`
}

func main() {
	opts := options.NewOptions()
	master.RunServer(opts)("hurricane")
}
