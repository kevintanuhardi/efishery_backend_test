package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var cfg Config

// Load reads in config file and ENV variables if set.
func Load(file string) (*Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
			fmt.Println(err.Error())
}

	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return &config, err
}
