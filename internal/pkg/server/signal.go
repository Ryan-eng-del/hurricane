package server

import (
	"os"
	"os/signal"
	"sync"
)


var setupOnceHandler sync.Once
var stop chan struct{}
var shoutDownHandler chan os.Signal


func SetupSignalHandler () <-chan struct{} {
	setupOnceHandler.Do(func(){
		shoutDownHandler = make(chan os.Signal, 2)
		signal.Notify(shoutDownHandler, shutdownSignals...)

		go func ()  {
			<- shoutDownHandler
			close(stop)
			<- shoutDownHandler
		  os.Exit(1)
		}()
	})

	return stop
}

func RequestShutdown() bool {
	if shoutDownHandler != nil {
		select {
		case shoutDownHandler <- shutdownSignals[0]:
			return true
		default:
		}
	}
	return false
}