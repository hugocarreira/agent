package project

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/hugocarreira/agentrc/cli/internal/config"
	"github.com/hugocarreira/agentrc/cli/internal/log"
)

func Link(name string) {
	projectDir := config.ProjectDir(name)
	agentsFile := filepath.Join(projectDir, "AGENTS.md")

	if _, err := os.Stat(agentsFile); os.IsNotExist(err) {
		log.L.Info("❌ No AGENTS.md found at " + agentsFile)
		log.L.Info("")
		log.L.Info("For new projects, use: agentrc init " + name)
		return
	}

	data, err := os.ReadFile(agentsFile)
	if err != nil {
		log.L.Info("❌ Failed to read AGENTS.md: " + err.Error())
		return
	}

	root := config.AgentRoot()
	if root == "" {
		log.L.Info("❌ Could not determine agentrc root. Set AGENT_ROOT env var.")
		return
	}

	ref := "READ " + root + "/AGENTS.md"
	if strings.Contains(string(data), ref) {
		log.L.Info("✅ AGENTS.md already references global config")
		return
	}

	backup := agentsFile + ".backup." + strconv.FormatInt(time.Now().Unix(), 10)
	os.Rename(agentsFile, backup)
	log.L.Info("📦 Backup created: " + backup)

	content := fmt.Sprintf("%s BEFORE ANYTHING.\n\n", ref)
	content += string(data)

	if err := os.WriteFile(agentsFile, []byte(content), 0644); err != nil {
		log.L.Info("❌ Failed to write AGENTS.md: " + err.Error())
		os.Rename(backup, agentsFile)
		return
	}

	log.L.Info("✅ Added global reference to: " + agentsFile)
}