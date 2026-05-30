package cmd

import (
	"github.com/hugocarreira/agentrc/cli/internal/verify"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Check harness setup and the current project pack",
	Long: `Verifies that:
- Config files exist for all installed agents
- Skills directories are linked
- RTK is available
- Current project state files are valid when present`,
	Run: runVerify,
}

func runVerify(cmd *cobra.Command, args []string) {
	verify.Run()
}
