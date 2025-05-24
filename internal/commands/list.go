package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/TakahashiShuuhei/todo/internal/utils"
	"github.com/spf13/cobra"
)

var (
	showArchived bool
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all tasks",
	Long:    `List all tasks in the todo list.`,
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

		// タブライターを初期化
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tStatus\tTitle\tDescription\tCreated At\tUpdated At")
		fmt.Fprintln(w, "--\t------\t-----\t-----------\t----------\t----------")

		// タスクを表示
		for _, todo := range *todos {
			// アーカイブされていないタスクのみを表示（showArchivedがfalseの場合）
			if !showArchived && todo.Archived {
				continue
			}

			status := " "
			if todo.Completed {
				status = "✓"
			}
			if todo.Archived {
				status = "🗄"
			}

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\n",
				todo.ID,
				status,
				todo.Title,
				todo.Description,
				todo.CreatedAt.Format("2006-01-02 15:04:05"),
				todo.UpdatedAt.Format("2006-01-02 15:04:05"),
			)
		}
		w.Flush()
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showArchived, "archived", "a", false, "Show archived tasks")
}
