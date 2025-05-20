package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/TakahashiShuuhei/todo/internal/models"
)

const (
	configDir  = ".todo"
	configFile = "config.json"
)

func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configDir, configFile), nil
}

func LoadConfig() (*models.Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	// 設定ファイルが存在しない場合はデフォルト設定を返す
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return models.DefaultConfig(), nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config models.Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *models.Config) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	// 設定ディレクトリが存在しない場合は作成
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
} 