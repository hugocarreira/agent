package cmd

import (
	"github.com/hugocarreira/agentrc/cli/internal/config"
	"github.com/hugocarreira/agentrc/cli/internal/log"
	"github.com/hugocarreira/agentrc/cli/internal/project"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Create AGENTS.md for a new project",
	Long: `Copies AGENTS.template.md to <project>/AGENTS.md.

For projects that don't have an AGENTS.md yet.
Backs up existing file before overwriting.`,
	Args: cobra.ExactArgs(1),
	Run: runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	root := config.AgentRoot()
	if root == "" {
		log.L.Info("❌ Could not find agentrc repo root.")
		log.L.Info("   Run from within the repo directory or set AGENT_ROOT.")
		return
	}
	project.New(root, args[0])
}
