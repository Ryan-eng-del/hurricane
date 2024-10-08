package main

import (
	"hurricane/internal/master"
	"math/rand"
	"time"

	_ "go.uber.org/automaxprocs"
)

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	master.NewApp("hurricane").Run()
}
