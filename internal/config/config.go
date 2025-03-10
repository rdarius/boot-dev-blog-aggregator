package config

import (
	"encoding/json"
	"os"
	"os/user"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	return cfg, nil

}

func (config Config) SetUser() error {
	usr, err := user.Current()

	if err != nil {
		return err
	}

	config.CurrentUserName = usr.Username

	err = write(config)
	if err != nil {
		return err
	}

	return nil

}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // For pretty printing
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
