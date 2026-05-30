package cmd

import (
	"github.com/hugocarreira/agentrc/cli/internal/config"
	"github.com/hugocarreira/agentrc/cli/internal/log"
	"github.com/hugocarreira/agentrc/cli/internal/project"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Create AGENTS.md, FEATURES.yaml, and PROGRESS.md for a project",
	Long: `Creates the default project pack in <project>:
- AGENTS.md
- FEATURES.yaml
- PROGRESS.md

Backs up existing generated files before overwriting.
If no project-name is given, runs in the current directory.`,
	Args: cobra.MaximumNArgs(1),
	Run:  runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	root := config.AgentRoot()
	if root == "" {
		log.L.Info("❌ Could not find agentrc repo root.")
		log.L.Info("   Run from within the repo directory or set AGENT_ROOT.")
		return
	}
	name := "."
	if len(args) > 0 {
		name = args[0]
	}
	project.New(root, name)
}
