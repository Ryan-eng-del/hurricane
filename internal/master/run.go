package master

import "github.com/Ryan-eng-del/hurricane/internal/master/config"

func Run(config *config.Config) error {
	server, err := createMasterApiServer(config)
	if err != nil {
		return err
	}
	return server.Run()
}
