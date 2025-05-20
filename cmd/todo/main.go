package main

import (
	"fmt"
	"os"

	"github.com/TakahashiShuuhei/todo/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a simple CLI tool for managing your tasks",
	Long: `A simple and efficient CLI tool for managing your daily tasks.
You can add, list, complete, and delete tasks easily.`,
}

func init() {
	commands.Init(rootCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
} 