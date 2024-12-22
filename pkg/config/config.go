package config

import (
    "encoding/json"
    "os"
)

// Config represents the application configuration
type Config struct {
    AurAPIURL string `json:"aur_api_url"`
}

// LoadConfig loads the configuration from a file
func LoadConfig() (*Config, error) {
    file, err := os.Open("config.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var cfg Config
    if err := json.NewDecoder(file).Decode(&cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}
