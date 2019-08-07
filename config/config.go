package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	defaultConfigPath = "/etc/simple-cache-service/"
)

var (
	configPath    = ""
	validBackends = []string{"redis", "bolt", "memory"}
)

// Config is the configuration for simple redis cache
type Config struct {
	Backend     string
	Compression string
	BoltOpts    BoltOpts
	RedisOpts   RedisOpts
}

// RedisOpts has options for redis
type RedisOpts struct {
	Address string
	Port    int
}

// BoltOpts has options for BoltDB
type BoltOpts struct {
}

// Get gets the config
func Get() (*Config, error) {
	configPath = getConfigPath()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}

	backend := viper.GetString("backend")
	if !validBackend(backend) {
		return nil, fmt.Errorf("%s is not a valid backend\n", backend)
	}

	return &Config{
		Backend:     backend,
		Compression: viper.GetString("compression"),
		RedisOpts: RedisOpts{
			Address: viper.GetString("redis_opts.address"),
			Port:    viper.GetInt("redis_opts.port"),
		},
	}, nil
}

func getConfigPath() string {
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = defaultConfigPath
	}
	return configPath
}

// validBackend will return true if backend is valid
func validBackend(backend string) bool {
	for _, be := range validBackends {
		if backend == be {
			return true
		}
	}
	return false
}
