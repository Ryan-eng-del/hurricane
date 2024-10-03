package master

import (
	"hurricane/internal/master/config"
	"hurricane/internal/pkg/server"
)

type MasterApiServer struct {
	APIServer *server.BaseApiServer
}

func (s *MasterApiServer) Run() error {
	return s.APIServer.Run()
}

func createMasterApiServer(config *config.Config) (*MasterApiServer, error) {
	apiServerConfig, err := buildApiServerConfig(config)
	if err != nil {
		return nil, err
	}
	apiServer, err := apiServerConfig.NewServer()

	if err != nil {
		return nil, err
	}

	return &MasterApiServer{
		APIServer: apiServer,
	}, nil
}

func buildApiServerConfig(config *config.Config) (serverConfig *server.Config, lastErr error) {
	serverConfig = server.NewConfig()
	if lastErr = config.GenericServerRunOptions.ApplyTo(serverConfig); lastErr != nil {
		return
	}

	if lastErr = config.SecureServing.ApplyTo(serverConfig); lastErr != nil {
		return
	}

	if lastErr = config.FeatureOptions.ApplyTo(serverConfig); lastErr != nil {
		return
	}

	if lastErr = config.InsecureServing.ApplyTo(serverConfig); lastErr != nil {
		return
	}
	return
}
