package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// TODO(TA)
	// Test getting config with no CONFIG_PATH
	// Test getting config with CONFIG_PATH
}

func TestGetConfigPath(t *testing.T) {
	os.Setenv("CONFIG_PATH", "")

	shouldBeDefault := getConfigPath()
	assert.Equal(t, defaultConfigPath, shouldBeDefault)

	testPath := "/tmp/simple-cache-service/"
	os.Setenv("CONFIG_PATH", testPath)

	shouldNotBeDefault := getConfigPath()
	assert.Equal(t, testPath, shouldNotBeDefault)
}

func TestValidBackend(t *testing.T) {
	valid := "redis"
	notValid := "not_redis"

	shouldBeValid := validBackend(valid)
	assert.True(t, shouldBeValid)

	shouldNotBeValid := validBackend(notValid)
	assert.False(t, shouldNotBeValid)
}
