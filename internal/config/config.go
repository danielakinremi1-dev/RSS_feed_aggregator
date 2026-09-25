package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Dburl    string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func (c Config) SetUser(username string) error {
	c.Username = username
	if err := write(c); err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err

	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	configData := Config{}
	err = json.Unmarshal(data, &configData)
	if err != nil {
		return Config{}, err
	}

	return configData, nil

}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := homePath + "/" + configFileName
	return configPath, nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	cfgData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, []byte(cfgData), 0644)
	if err != nil {
		return err
	}

	return nil
}

const configFileName = ".gatorconfig.json"
