package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getGlobalConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, GLOBAL_CONFIG_FILE), nil
}

func LoadGlobalConfig() Config {
	cfg := GetDefaultConfig()
	path, err := getGlobalConfigPath()
	if err != nil {
		fmt.Printf("Warning: could not get home directory: %v\n", err)
		return cfg
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return cfg
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Warning: reading global config: %v\n", err)
		return cfg
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Printf("Warning: parsing global config: %v\n", err)
	}

	return cfg
}

func LoadProjectConfig() Config {
	path := filepath.Join(".", PROJECT_CONFIG_FILE)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return GetDefaultConfig()
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return GetDefaultConfig()
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return GetDefaultConfig()
	}
	return cfg
}

func SaveGlobalConfig(cfg Config) error {
	path, err := getGlobalConfigPath()
	if err != nil {
		return err
	}
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(path, data, 0644)
}

func SaveProjectConfig(cfg Config) error {
	path := filepath.Join(".", PROJECT_CONFIG_FILE)
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(path, data, 0644)
}
