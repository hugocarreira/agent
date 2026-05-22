package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hugocarreira/agentrc/cli/internal/config"
	"github.com/hugocarreira/agentrc/cli/internal/harness"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show overview of agents, config, and commands",
	Long:  `Displays a summary of installed agents, their config status, RTK availability, and the repo root.`,
	Run: runStatus,
}

func runStatus(cmd *cobra.Command, args []string) {
	root := config.AgentRoot()
	rtk := strings.TrimSpace(harness.RTKVersion())
	agents := harness.AllAgents()

	fmt.Println()
	fmt.Println("  agentrc — AI Agent Setup Manager")
	fmt.Println("  " + strings.Repeat("─", 36))
	fmt.Println()

	if root != "" {
		fmt.Printf("  Repo:     %s\n", root)
	} else {
		fmt.Println("  Repo:     not found (run from repo dir or set AGENT_ROOT)")
	}

	if rtk != "" {
		fmt.Printf("  RTK:      %s\n", rtk)
	} else {
		fmt.Println("  RTK:      not installed (optional)")
	}

	fmt.Println("  Commands: setup | init <name> | link <name> | verify")
	fmt.Println()

	for _, a := range agents {
		if !a.Installed {
			fmt.Printf("  %-18s %s\n", a.Name, "—")
			continue
		}

		cfgOk := fileExists(filepath.Join(a.Dir, a.ConfigFile))
		skillsOk := symlinkExists(filepath.Join(a.Dir, "skills"))

		parts := []string{}
		if cfgOk {
			parts = append(parts, "✓ cfg")
		} else {
			parts = append(parts, "✗ cfg")
		}
		if skillsOk {
			parts = append(parts, "✓ skills")
		} else {
			parts = append(parts, "✗ skills")
		}

		if a.Name == "OpenCode" {
			pluginsOk := symlinkExists(filepath.Join(a.Dir, "plugins"))
			if pluginsOk {
				parts = append(parts, "✓ plugins")
			} else {
				parts = append(parts, "✗ plugins")
			}
		}

		fmt.Printf("  %-18s installed  %s\n", a.Name, strings.Join(parts, "  "))
	}
	fmt.Println()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func symlinkExists(path string) bool {
	fi, err := os.Lstat(path)
	return err == nil && fi.Mode()&os.ModeSymlink != 0
}
