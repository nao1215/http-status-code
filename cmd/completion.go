package cmd

import (
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Create shell completion files (bash, fish, zsh)",
	Long: `Create shell completion files (bash, fish, zsh) for the hsc command
if it is not already on the system`,
	Run: func(cmd *cobra.Command, args []string) {
		deployShellCompletionFileIfNeeded(rootCmd)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
