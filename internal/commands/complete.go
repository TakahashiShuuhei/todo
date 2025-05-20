package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/TakahashiShuuhei/todo/internal/utils"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:     "complete [id]",
	Aliases: []string{"c"},
	Short:   "Complete a task",
	Long:    `Complete a task by ID.`,
	Args:    cobra.ExactArgs(1),
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

		if len(*todos) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		// IDを解析
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID: %s", args[0])
		}

		// タスクを探して完了状態をトグル
		for i, todo := range *todos {
			if todo.ID == id {
				(*todos)[i].Completed = !todo.Completed
				(*todos)[i].UpdatedAt = time.Now()
				if err := utils.SaveTodos(todoPath, todos); err != nil {
					return fmt.Errorf("failed to save todos: %w", err)
				}
				status := "completed"
				if !(*todos)[i].Completed {
					status = "uncompleted"
				}
				fmt.Printf("Task %d %s\n", id, status)
				return nil
			}
		}
		return fmt.Errorf("task with ID %d not found", id)
	},
}
