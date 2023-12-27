/**
Note: run this test from main folder, run: go test
*/

package config

import (
	"testing"
)

var (
	config *Config
	err    error
)

func TestLoadConfig(t *testing.T) {
	config, err = LoadConfig()
	if err != nil {
		t.Errorf("Failed to load config: %v\n", err)
		return
	}
	t.Log("Loaded config")
}

// test get env value
func TestGetEnv(t *testing.T) {
	if config.Port == "" {
		t.Errorf("Failed to get env: APP_PORT\n")
		return
	}

	if config.DB.Host == "" {
		t.Errorf("Failed to get env: DB_HOST\n")
		return
	}

	if config.DB.Database == "" {
		t.Errorf("Failed to get env: DB_NAME\n")
		return
	}

	if config.DB.Username == "" {
		t.Errorf("Failed to get env: DB_USERNAME\n")
	}
}
