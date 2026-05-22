package cmd

import (
	"github.com/hugocarreira/agentrc/cli/internal/project"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link [project-dir]",
	Short: "Add global reference to existing project",
	Long: `Prepends a global AGENTS.md reference to existing project config.

For projects that already have their own AGENTS.md.
If no project-dir is given, uses the current directory.`,
	Args: cobra.MaximumNArgs(1),
	Run: runLink,
}

func runLink(cmd *cobra.Command, args []string) {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}
	project.Link(dir)
}
