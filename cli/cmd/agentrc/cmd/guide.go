package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var guideCmd = &cobra.Command{
	Use:   "guide",
	Short: "Show a friendly overview of agentrc",
	Args:  cobra.NoArgs,
	Run:   runGuide,
}

func runGuide(cmd *cobra.Command, args []string) {
	fmt.Println()
	fmt.Println("  agentrc — AI Agent Setup Manager")
	fmt.Println("  " + strings.Repeat("─", 36))
	fmt.Println()
	fmt.Println("  Central config for 6 AI coding agents.")
	fmt.Println("  One repo to rule them all.")
	fmt.Println()
	fmt.Println("  ── Quick Start ──")
	fmt.Println()
	fmt.Println("    agentrc setup     Link config + skills to all installed agents")
	fmt.Println("    agentrc verify    Check setup plus current project state files")
	fmt.Println("    agentrc status    Overview of agents, config, commands")
	fmt.Println()
	fmt.Println("  ── Per-Project ──")
	fmt.Println()
	fmt.Println("    agentrc init <n>  Create AGENTS.md + FEATURES.yaml + PROGRESS.md")
	fmt.Println("    agentrc link <n>  Add global reference to existing project")
	fmt.Println()
	fmt.Println("  ── Tips ──")
	fmt.Println()
	fmt.Println("    • Run commands from inside the agentrc repo")
	fmt.Println("    • Or set AGENT_ROOT to point to it")
	fmt.Println("    • Edit AGENTS.md / RTK.md in the repo, then re-run setup")
	fmt.Println("    • Use agentrc <cmd> --help for detailed flags")
	fmt.Println()
}
