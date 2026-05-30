package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agentrc",
	Short: "AI agent harness setup & project management",
	Long: `agentrc — manage 6 AI coding agents (OpenCode, Codex, Claude, Copilot, Hermes, Gemini).

Setup configs, init project packs, link existing projects, verify setup, check status.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(linkCmd)
	rootCmd.AddCommand(verifyCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(guideCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
