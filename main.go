// Copyright 2024 Benjamin Lee <cyan0908@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/marmotedu/errors"
)

type Info struct {
	GitCommit  string   `json:"git_commit"`
	GitVersion int      `json:"git_version"`
	GitArr     []string `json:"git_"`
}
type PathError struct {
	Name string
}

type PathError1 struct {
}

var Err2 = errors.New("Error2")
var Err1 = errors.New("Error1")

var Err3 = &PathError{}
var Err4 = &PathError{"sdf"}

var Err5 = &PathError1{}

func (*PathError) Error() string {
	return "13"
}

func (*PathError) Err3Switch() string {
	return "1133"
}

func (*PathError1) Error() string {
	return "13"
}

func (*PathError1) Err3Switch() string {
	return "1133"
}

func Error2() error {
	return Err1
}

func Error1() error {
	if err := Error2(); err != nil {
		return errors.Wrap(err, Err2.Error())
	}
	return nil
}

func main() {
	if err := Error1(); err != nil {
		log.Println(err)
	}
}
