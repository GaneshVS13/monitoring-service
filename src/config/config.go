package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Config represents the config entity
type Config struct {
	ServiceConfig  ServiceConfig `json:"ServiceConfig"`
	MonitoringURLs []string      `json:"URLs"`
}

// ServiceConfig is an entity for Service configuration
type ServiceConfig struct {
	ListenURL string `json:"ListenURL"`
}

// Load loads config once
func Load(configPath string) (Config, error) {
	var cfg Config

	err := readFromJSON(configPath, &cfg)
	if err != nil {
		return cfg, errors.New("configuration not found")
	}

	return cfg, nil
}

// readFromJSON reads config data from JSON-file
func readFromJSON(configFilePath string, cfg *Config) error {
	contents, err := ioutil.ReadFile(filepath.Clean(configFilePath))
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(cfg)
	}
	if err != nil {
		log.Printf("failed to read (%s). err: %v\n", configFilePath, err)
		return err
	}

	log.Printf("successfully read (%s)\n", configFilePath)

	return nil
}
