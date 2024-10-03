package options

import "hurricane/internal/pkg/server"

// ServerRunOptions contains the options while running a generic api server.
type ServerRunOptions struct {
	Mode         string   `json:"mode"        mapstructure:"mode"`
	EnableHealth bool     `json:"healthz"     mapstructure:"healthz"`
	Middlewares  []string `json:"middlewares" mapstructure:"middlewares"`
}

// NewServerRunOptions creates a new ServerRunOptions object with default parameters.
func NewServerRunOptions() *ServerRunOptions {
	defaults := server.NewConfig()

	return &ServerRunOptions{
		Mode:         defaults.Mode,
		EnableHealth: defaults.EnableHealth,
		Middlewares:  defaults.Middlewares,
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *ServerRunOptions) ApplyTo(c *server.Config) error {
	c.Mode = s.Mode
	c.EnableHealth = s.EnableHealth
	c.Middlewares = s.Middlewares

	return nil
}
