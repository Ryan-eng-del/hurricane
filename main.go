// Copyright 2024 Benjamin Lee <cyan0908@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
