package main

import (
	"math/rand"
	"time"

	"github.com/Ryan-eng-del/hurricane/internal/master"

	_ "github.com/Ryan-eng-del/hurricane/third_party/forked/automaxprocs"
)

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	master.NewApp("hurricane").Run()
}
