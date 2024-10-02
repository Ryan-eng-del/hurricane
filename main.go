package main

import (
	"hurricane/pkg/log"
)

func main() {
	log.NewWithOptions(log.WithEnableFile(), log.WithEnableColor(), log.WithDebugMode())
}
