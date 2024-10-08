package main

import (
	"log"
	"strings"
)

type Info struct {
	GitCommit  string   `json:"git_commit"`
	GitVersion int      `json:"git_version"`
	GitArr     []string `json:"git_"`
}

func main() {
	log.Println(strings.Index("12345", "123"))
}
