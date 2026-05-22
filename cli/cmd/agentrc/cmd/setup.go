package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func promptRoot() string {
	detected := config.AgentRoot()
	scanner := bufio.NewScanner(os.Stdin)

	if detected != "" {
		fmt.Printf("Detected agentrc repo at: %s\n", detected)
		fmt.Printf("Use this path? [Y/n]: ")
		scanner.Scan()
		answer := strings.TrimSpace(scanner.Text())
		if answer == "" || strings.EqualFold(answer, "y") || strings.EqualFold(answer, "yes") {
			return detected
		}
	}

	fmt.Print("Enter the full path to your agentrc repo: ")
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func runSetup(cmd *cobra.Command, args []string) {
	root := config.AgentRoot()
	if root == "" {
		log.L.Info("AGENT_ROOT not set.")
		root = promptRoot()
		if root == "" {
			log.L.Info("No path provided. Aborting.")
			return
		}
		if err := config.SaveAgentRootToShell(root); err != nil {
			log.L.Info("⚠️  Could not save AGENT_ROOT to shell config: " + err.Error())
			log.L.Info("   Set it manually: export AGENT_ROOT=" + root)
		} else {
			log.L.Info("✅ AGENT_ROOT saved to " + config.ShellConfigFile())
			log.L.Info("   Restart your terminal or run: source " + config.ShellConfigFile())
		}
	}
	harness.Setup(root)
}
