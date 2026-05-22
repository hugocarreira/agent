package cmd

import (
	"github.com/hugocarreira/agentrc/cli/internal/project"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link <project-name>",
	Short: "Add global reference to existing project",
	Long: `Prepends a global AGENTS.md reference to existing project config.

For projects that already have their own AGENTS.md.`,
	Args: cobra.ExactArgs(1),
	Run: runLink,
}

func runLink(cmd *cobra.Command, args []string) {
	project.Link(args[0])
}
