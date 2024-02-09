package util

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type (
	Cmds     []string
	GroupMap map[string]Cmds
	CmdsMap  map[string]GroupMap
)

const (
	DIR_NAME       = ".mycmd"
	JSON_FILE_NAME = "cmd.json"
)

func InitJson() {
	jsonFilePath, err := GetJsonFilePath()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
		initialData := map[string]interface{}{
			"Linux": map[string][]string{
				"list": {"ls __arg__"},
			},
			"Windows": map[string][]string{
				"list": {"dir __arg__"},
			},
		}

		jsonData, err := json.MarshalIndent(initialData, "", "  ")
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(jsonFilePath, jsonData, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func GetJsonFilePath() (jsonFilePath string, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, DIR_NAME)

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.Mkdir(configDir, 0755); err != nil {
			return "", err
		}
	}

	return filepath.Join(configDir, JSON_FILE_NAME), nil
}

func GetCmdsMap() (cmdsMap CmdsMap, err error) {
	jsonFilePath, err := GetJsonFilePath()
	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &cmdsMap); err != nil {
		return nil, err
	}

	return cmdsMap, nil
}
