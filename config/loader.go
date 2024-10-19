package config

import (
	"fmt"
	"lockbox/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
)

// Pointer to hold the loaded config
var conf *Config

// GetConfig returns the cached config, loading it if it hasn't been loaded yet
func GetConfig() *Config {
	if conf != nil { // If config is already loaded, return it
		return conf
	}

	conf, err := loadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// If config hasn't been loaded yet, load it
	return conf
}

func loadConfig() (*Config, error) {
	// Get the user's home directory
	homeDir, err := utils.Dir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return nil, err
	}

	// Initialize filePath
	var filePath string
	// Determine the path based on the operating system
	if os.PathSeparator == '\\' {
		filePath = filepath.Join(homeDir, "AppData", "Local", "lockbox", "config.toml") // Windows path
	} else {
		filePath = filepath.Join(homeDir, ".config", "lockbox", "config.toml") // Unix-like path
	}

	// Read the TOML file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// Parse the TOML data
	conf = &Config{} // Initialize config pointer
	if err := toml.Unmarshal(data, conf); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	if conf.PasswordsStorePath == "" {
		if os.PathSeparator == '\\' {
			conf.PasswordsStorePath = filepath.Join(homeDir, "AppData", "Roaming", "lockbox") // Unix-like path
		} else {
			conf.PasswordsStorePath = filepath.Join(homeDir, ".local", "share", "lockbox") // Unix-like path
		}
	} else {
		// Check if passwords_store_path starts with ~ and replace it with the home directory
		if strings.HasPrefix(conf.PasswordsStorePath, "~") {
			conf.PasswordsStorePath = filepath.Join(homeDir, conf.PasswordsStorePath[1:]) // Replace ~ with homeDir
		}
	}

	return conf, nil
}
