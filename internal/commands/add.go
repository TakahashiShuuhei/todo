package commands

import (
	"fmt"
	"time"

	"github.com/TakahashiShuuhei/todo/internal/models"
	"github.com/TakahashiShuuhei/todo/internal/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [title] [description]",
	Aliases: []string{"a"},
	Short:   "Add a new task",
	Long:    `Add a new task with a title and optional description.`,
	Args:    cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Todoファイルを探す
		todoPath, err := utils.FindTodoFile()
		if err != nil {
			return fmt.Errorf("failed to find todo file: %w", err)
		}
		if todoPath == "" {
			return fmt.Errorf("todo file not found. please run 'todo init' first")
		}

		// 既存のタスクを読み込む
		todos, err := utils.LoadTodos(todoPath)
		if err != nil {
			return fmt.Errorf("failed to load todos: %w", err)
		}

		// 新しいタスクを作成
		now := time.Now()
		newTodo := models.Todo{
			ID:        len(*todos) + 1,
			Title:     args[0],
			Completed: false,
			CreatedAt: now,
			UpdatedAt: now,
		}

		// 説明が指定されている場合は追加
		if len(args) > 1 {
			newTodo.Description = args[1]
		}

		// タスクを追加
		*todos = append(*todos, newTodo)

		// 保存
		if err := utils.SaveTodos(todoPath, todos); err != nil {
			return fmt.Errorf("failed to save todos: %w", err)
		}

		fmt.Printf("Added task: %s\n", newTodo.Title)
		return nil
	},
}
