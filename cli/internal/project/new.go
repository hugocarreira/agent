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

type templateFile struct {
	Template string
	Target   string
	Build    func(root string, tmpl []byte) []byte
}

func managedTemplateFiles() []templateFile {
	return []templateFile{
		{
			Template: "AGENTS.template.md",
			Target:   "AGENTS.md",
			Build:    buildAgentsFile,
		},
		{
			Template: "FEATURES.template.yaml",
			Target:   "FEATURES.yaml",
		},
		{
			Template: "PROGRESS.template.md",
			Target:   "PROGRESS.md",
		},
	}
}

func New(root, name string) {
	projectDir := config.ProjectDir(name)

	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		log.L.Info("❌ Project '" + name + "' not found at " + projectDir)
		log.L.Info("")
		log.L.Info("Create it first: mkdir -p " + projectDir)
		return
	}

	files := managedTemplateFiles()

	rendered, err := loadTemplates(root, files)
	if err != nil {
		log.L.Info("❌ " + err.Error())
		return
	}

	existing := existingManagedFiles(projectDir, files)
	if len(existing) > 0 {
		log.L.Info("⚠️  Project pack already exists in " + projectDir)
		log.L.Info("Overwrite generated files? [y/N]: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() != "y" && scanner.Text() != "Y" {
			log.L.Info("Cancelled.")
			return
		}
		backupManagedFiles(existing)
	}

	for target, data := range rendered {
		targetPath := filepath.Join(projectDir, target)
		if err := os.WriteFile(targetPath, data, 0644); err != nil {
			log.L.Info("❌ Failed to write " + target + ": " + err.Error())
			return
		}
		log.L.Info("✅ Created: " + targetPath)
	}

	log.L.Info("")
	log.L.Info("➡️  Next instruction for your AI agent:")
	log.L.Info("")
	log.L.Info(initInstruction())
}

func buildAgentsFile(root string, tmpl []byte) []byte {
	ref := fmt.Sprintf("READ %s/AGENTS.md BEFORE ANYTHING.\n\n", root)
	instruction := "> **AI:** Generated from template. Scan project codebase, update Stack/Commands/Out of Scope, replace starter feature entries, and remove this line when done.\n\n---\n\n"
	data := []byte(ref + instruction)
	return append(data, tmpl...)
}

func loadTemplates(root string, files []templateFile) (map[string][]byte, error) {
	rendered := make(map[string][]byte, len(files))

	for _, file := range files {
		templatePath := filepath.Join(root, file.Template)
		tmpl, err := os.ReadFile(templatePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read template %s: %w", templatePath, err)
		}

		data := tmpl
		if file.Build != nil {
			data = file.Build(root, tmpl)
		}
		rendered[file.Target] = data
	}

	return rendered, nil
}

func existingManagedFiles(projectDir string, files []templateFile) []string {
	var existing []string

	for _, file := range files {
		targetPath := filepath.Join(projectDir, file.Target)
		if _, err := os.Stat(targetPath); err == nil {
			existing = append(existing, targetPath)
		}
	}

	return existing
}

func backupManagedFiles(paths []string) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	for _, path := range paths {
		backup := path + ".backup." + timestamp
		if err := os.Rename(path, backup); err != nil {
			log.L.Info("⚠️  Failed to back up " + path + ": " + err.Error())
			continue
		}
		log.L.Info("📦 Backup created: " + backup)
	}
}

func initInstruction() string {
	return `Read AGENTS.md first. Then inspect this repo and finish the project pack from evidence:

1. Update AGENTS.md so Stack, Commands, Out of Scope, and Project Notes match the real codebase.
2. Replace the starter entries in FEATURES.yaml with a short list of concrete current features. For each feature, set behavior, done_when, and state.
3. Rewrite PROGRESS.md so it reflects the current session state, blockers, and exact next step.

Use code, docs, package scripts, and config files as sources of truth. Do not invent product strategy, roadmap, or requirements. If something is unclear, leave a TODO or ask the user.`
}
