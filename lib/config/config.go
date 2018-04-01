package config

import (
	"github.com/spf13/viper"
)

const (
	// EnvPrefix to minimize the chance to mix with some settings for other apps
	EnvPrefix = "FIB"

	// EnvDefaultPort is the default Port.
	EnvDefaultPort = 8080

	// EnvDefaultCacheNumbers is the default cache numbers
	EnvDefaultCacheNumbers = 10000

	// EnvDefaultMaximumNumbers is the default maximum numbers we supported
	EnvDefaultMaximumNumbers = 20000

	// EnvDefaultLogLevel is the default log level
	EnvDefaultLogLevel = "debug"
)

// Config represents the config object for Server.
type Config struct {
	// Port represents the Listening port of the http server.
	Port uint

	// LogLevel represents the log level used in the application.
	LogLevel string

	// CachedNumbers represents the fibonacci numbers to cache.
	CachedNumbers uint64

	// MaximumNumbers represents the maximum return fibonacci numbers we support
	MaximumNumbers uint64

	// Version
	Version string

	// Build Date
	BuildDate string
}

// Build is used to create a Config instance from configuraion file and environment
// variables.
func Build() (*Config, error) {
	v := viper.GetViper()
	v.SetEnvPrefix(EnvPrefix)

	v.SetDefault("Port", EnvDefaultPort)
	v.SetDefault("LogLevel", EnvDefaultLogLevel)
	v.SetDefault("CachedNumbers", EnvDefaultCacheNumbers)
	v.SetDefault("MaximumNumbers", EnvDefaultMaximumNumbers)

	v.BindEnv("Port")
	v.BindEnv("LogLevel")
	v.BindEnv("CachedNumbers")
	v.BindEnv("MaximumNumbers")

	return &Config{
		Port:           uint(v.GetInt("Port")),
		LogLevel:       v.GetString("LogLevel"),
		CachedNumbers:  uint64(v.GetInt("CachedNumbers")),
		MaximumNumbers: uint64(v.GetInt("MaximumNumbers")),
		Version:        viper.Get("VERSION").(string),
		BuildDate:      viper.Get("BUILDDATE").(string),
	}, nil
}
