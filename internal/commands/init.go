package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/TakahashiShuuhei/todo/internal/models"
	"github.com/TakahashiShuuhei/todo/internal/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new todo file in the current directory",
	Long:  `Initialize a new todo file in the current directory. The file name is configurable in ~/.todo/config.json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := utils.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		// 現在のディレクトリに.todoファイルが既に存在するか確認
		todoPath := filepath.Join(".", config.TodoFileName)
		if _, err := os.Stat(todoPath); err == nil {
			return fmt.Errorf("todo file already exists at %s", todoPath)
		}

		// 空のTodoListを作成
		todos := models.TodoList{}

		// ファイルに保存
		if err := utils.SaveTodos(todoPath, &todos); err != nil {
			return fmt.Errorf("failed to create todo file: %w", err)
		}

		fmt.Printf("Initialized todo file at %s\n", todoPath)
		return nil
	},
}

// このファイルではコマンド登録は行わない 