package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Source       string   `json:"source" validate:"required,udp_addr"`
	Destinations []string `json:"destinations" validate:"required,min=1,dive,udp_addr"`
}

func (c *Config) String() string {
	return fmt.Sprintf("Source: %s, Destinations: %s", c.Source, strings.Join(c.Destinations, ", "))
}

func getConfiguration(flags *Flags) (*Config, error) {
	config, err := readConfigFile(*flags.ConfigFilePath)
	if err != nil {
		return nil, err
	}

	err = config.Validate()
	if err != nil {
		return nil, err
	}

	Log.Info.Printf("Configuration: %s", config.String())

	return config, nil
}

func readConfigFile(filePath string) (*Config, error) {
	Log.Debug.Printf("Reading config file '%s'", filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var out Config

	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Config) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
