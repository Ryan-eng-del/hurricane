package main

import "fmt"

type Info struct {
	GitCommit  string   `json:"git_commit"`
	GitVersion int      `json:"git_version"`
	GitArr     []string `json:"git_"`
}

func main() {
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
}
