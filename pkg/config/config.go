package config

import (
	"encoding/json"
	"os"
)

//Config struct
type Config struct {
	Port string `json:"port"`
}

//LoadConfig : read json file and load the info into Config struct
func LoadConfig(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil

}
