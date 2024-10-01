package main

import (
	"hurricane/pkg/log"
)

func main() {

	log.NewWithOptions(log.WithEnableFile(), log.WithDebugMode())
	log.Log.Info("This is the info")
	log.Log.Debug("This is the debug")
	log.Log.Error("This is the error")
	log.Log.Warn("This is the warn")
}
