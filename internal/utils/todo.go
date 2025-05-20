package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/TakahashiShuuhei/todo/internal/models"
)

func FindTodoFile() (string, error) {
	config, err := LoadConfig()
	if err != nil {
		return "", err
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		todoPath := filepath.Join(dir, config.TodoFileName)
		if _, err := os.Stat(todoPath); err == nil {
			return todoPath, nil
		}

		// 親ディレクトリに移動
		parent := filepath.Dir(dir)
		if parent == dir {
			// ルートディレクトリに到達
			return "", nil
		}
		dir = parent
	}
}

func LoadTodos(todoPath string) (*models.TodoList, error) {
	data, err := os.ReadFile(todoPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &models.TodoList{}, nil
		}
		return nil, err
	}

	var todos models.TodoList
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return &todos, nil
}

func SaveTodos(todoPath string, todos *models.TodoList) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(todoPath, data, 0644)
} 