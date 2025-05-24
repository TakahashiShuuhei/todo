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
	RunE: func(cmd *cobra.Command, args []string) error {
		// Todoファイルを探す
		todoPath, err := utils.FindTodoFile()
		if err != nil {
			return fmt.Errorf("failed to find todo file: %w", err)
		}
		if todoPath == "" {
			return fmt.Errorf("todo file not found. please run 'todo init' first")
		}

		// タスクを読み込む
		todos, err := utils.LoadTodos(todoPath)
		if err != nil {
			return fmt.Errorf("failed to load todos: %w", err)
		}

		// 新しいタスクを作成
		newTodo := models.Todo{
			ID:          len(*todos) + 1,
			Title:       args[0],
			Description: "",
			Completed:   false,
			Archived:    false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// 説明が指定されている場合は追加
		if len(args) > 1 {
			newTodo.Description = args[1]
		}

		// タスクを追加
		*todos = append(*todos, newTodo)

		// タスクを保存
		if err := utils.SaveTodos(todoPath, todos); err != nil {
			return fmt.Errorf("failed to save todos: %w", err)
		}

		fmt.Printf("Added task: %s\n", newTodo.Title)
		fmt.Println() // 空行を追加
		PrintTodos(todos, false)
		return nil
	},
}

func init() {
	addCmd.Args = cobra.RangeArgs(1, 2)
}
