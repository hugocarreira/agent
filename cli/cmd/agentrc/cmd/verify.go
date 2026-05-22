package cmd

import (
	"github.com/hugocarreira/agentrc/cli/internal/verify"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Check config and skills setup",
	Long: `Verifies that:
- Config files exist for all installed agents
- Skills directories are linked
- RTK is available`,
	Run: runVerify,
}

func runVerify(cmd *cobra.Command, args []string) {
	verify.Run()
}
