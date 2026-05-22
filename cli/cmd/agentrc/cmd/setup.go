package cmd

import (
	"github.com/hugocarreira/agentrc/cli/internal/config"
	"github.com/hugocarreira/agentrc/cli/internal/log"
	"github.com/hugocarreira/agentrc/cli/internal/harness"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Global config setup (run once)",
	Long: `Writes config files and creates skill/plugin symlinks
so all AI agents share the same AGENTS.md, skills, and RTK config.

Run once after cloning the agentrc repo.`,
	Run: runSetup,
}

func runSetup(cmd *cobra.Command, args []string) {
	root := config.AgentRoot()
	if root == "" {
		log.L.Info("❌ Could not find agentrc repo root.")
		log.L.Info("   Run from within the repo directory or set AGENT_ROOT.")
		return
	}
	harness.Setup(root)
}
