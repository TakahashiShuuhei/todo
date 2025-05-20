package commands

import "github.com/spf13/cobra"

// Init initializes all commands
func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
}
