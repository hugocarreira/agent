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
	tmpl, err := os.ReadFile(templatePath)
	if err != nil {
		log.L.Info("❌ Failed to read template: " + err.Error())
		return
	}

	instruction := "> **AI:** Generated from template. Scan project codebase, then update the Stack and Commands sections below to match the actual setup. Delete this line when done.\n\n---\n\n"
	data := []byte(ref + instruction)
	data = append(data, tmpl...)

	if err := os.WriteFile(agentsFile, data, 0644); err != nil {
		log.L.Info("❌ Failed to write AGENTS.md: " + err.Error())
		return
	}

	log.L.Info("✅ Created: " + agentsFile)
	log.L.Info("")
	log.L.Info("➡️  Paste this into your AI agent:")
	log.L.Info("")
	log.L.Info("READ AGENTS.md AND USE AS BASE FOR THIS PROJECT. SCAN THE CODEBASE AND UPDATE Stack/Commands SECTIONS TO MATCH ACTUAL SETUP, THEN REMOVE THE INSTRUCTION LINE.")
}