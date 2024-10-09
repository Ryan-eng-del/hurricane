package config

import "github.com/Ryan-eng-del/hurricane/internal/master/options"

// Config is the running configuration structure of the hurricane service.
type Config struct {
	*options.Options
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given hurricane pump command line or configuration file option.
func CreateConfigFromOptions(options *options.Options) (*Config, error) {
	return &Config{options}, nil
}
