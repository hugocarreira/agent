package project

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/hugocarreira/agentrc/cli/internal/config"
	"github.com/hugocarreira/agentrc/cli/internal/log"
)

func New(root, name string) {
	projectDir := config.ProjectDir(name)
	agentsFile := filepath.Join(projectDir, "AGENTS.md")

	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		log.L.Info("❌ Project '" + name + "' not found at " + projectDir)
		log.L.Info("")
		log.L.Info("Create it first: mkdir -p " + projectDir)
		return
	}

	if _, err := os.Stat(agentsFile); err == nil {
		log.L.Info("⚠️  AGENTS.md already exists at " + agentsFile)
		log.L.Info("Overwrite? [y/N]: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() != "y" && scanner.Text() != "Y" {
			log.L.Info("Cancelled.")
			return
		}
		backup := agentsFile + ".backup." + strconv.FormatInt(time.Now().Unix(), 10)
		os.Rename(agentsFile, backup)
		log.L.Info("📦 Backup created: " + backup)
	}

	ref := fmt.Sprintf("READ %s/AGENTS.md BEFORE ANYTHING.\n\n", root)
	templatePath := filepath.Join(root, "AGENTS.template.md")
	data, err := os.ReadFile(templatePath)
	if err != nil {
		log.L.Info("❌ Failed to read template: " + err.Error())
		return
	}
	data = append([]byte(ref), data...)

	if err := os.WriteFile(agentsFile, data, 0644); err != nil {
		log.L.Info("❌ Failed to write AGENTS.md: " + err.Error())
		return
	}

	log.L.Info("✅ Created: " + agentsFile)
	log.L.Info("")
	log.L.Info("Next steps:")
	log.L.Info("  cd " + projectDir)
	log.L.Info("  # Edit AGENTS.md:")
	log.L.Info("  #   1. Choose stack option - delete others")
	log.L.Info("  #   2. Add project-specific commands")
	log.L.Info("  #   3. Keep it SHORT (50-100 lines)")
}