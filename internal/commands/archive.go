package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/TakahashiShuuhei/todo/internal/utils"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:     "archive [id|all]",
	Aliases: []string{"ar"},
	Short:   "Archive a task",
	Long:    `Archive a task by ID, all tasks, or interactively select tasks to archive.`,
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

		// "all" の場合は一括アーカイブのみ
		if len(args) > 0 && strings.ToLower(args[0]) == "all" {
			for i := range *todos {
				if !(*todos)[i].Archived {
					(*todos)[i].Archived = true
					(*todos)[i].UpdatedAt = time.Now()
				}
			}
			if err := utils.SaveTodos(todoPath, todos); err != nil {
				return fmt.Errorf("failed to save todos: %w", err)
			}
			fmt.Println("All tasks archived!")
			return nil
		}

		// IDが指定されている場合は直接アーカイブ
		if len(args) > 0 {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid task ID: %s", args[0])
			}

			for i, todo := range *todos {
				if todo.ID == id {
					(*todos)[i].Archived = !todo.Archived
					(*todos)[i].UpdatedAt = time.Now()
					if err := utils.SaveTodos(todoPath, todos); err != nil {
						return fmt.Errorf("failed to save todos: %w", err)
					}
					status := "archived"
					if !(*todos)[i].Archived {
						status = "unarchived"
					}
					fmt.Printf("Task %d %s\n", id, status)
					return nil
				}
			}
			return fmt.Errorf("task with ID %d not found", id)
		}

		// インタラクティブな選択UI
		var selected []int
		options := make([]huh.Option[int], len(*todos))
		for i, todo := range *todos {
			status := " "
			if todo.Archived {
				status = "🗄"
			}
			options[i] = huh.NewOption(fmt.Sprintf("[%d] %s %s", todo.ID, status, todo.Title), i)
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[int]().
					Title("Select tasks to toggle archive status").
					Options(options...).
					Value(&selected),
			),
		)

		if err := form.Run(); err != nil {
			return fmt.Errorf("failed to run form: %w", err)
		}

		// 選択されたタスクのアーカイブ状態をトグル
		for _, idx := range selected {
			(*todos)[idx].Archived = !(*todos)[idx].Archived
			(*todos)[idx].UpdatedAt = time.Now()
		}

		if err := utils.SaveTodos(todoPath, todos); err != nil {
			return fmt.Errorf("failed to save todos: %w", err)
		}

		fmt.Println("Tasks updated successfully!")
		return nil
	},
}
