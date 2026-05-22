package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AgentRoot() string {
	if env := os.Getenv("AGENT_ROOT"); env != "" {
		return env
	}
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		marker := filepath.Join(dir, "cli")
		if fi, err := os.Stat(marker); err == nil && fi.IsDir() {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

func ProjectDir(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	wd, err := os.Getwd()
	if err != nil {
		return path
	}
	return filepath.Join(wd, path)
}

func ShellConfigFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	shell := os.Getenv("SHELL")
	candidates := []string{}

	if strings.Contains(shell, "zsh") {
		candidates = append(candidates, filepath.Join(home, ".zshrc"))
	} else if strings.Contains(shell, "bash") {
		candidates = append(candidates, filepath.Join(home, ".bashrc"))
	}
	candidates = append(candidates, filepath.Join(home, ".profile"))

	for _, c := range candidates {
		if fi, err := os.Stat(c); err == nil && !fi.IsDir() {
			return c
		}
	}
	if len(candidates) > 0 {
		return candidates[0]
	}
	return ""
}

func SaveAgentRootToShell(root string) error {
	cfg := ShellConfigFile()
	if cfg == "" {
		return fmt.Errorf("could not find shell config file (.zshrc/.bashrc/.profile)")
	}
	line := fmt.Sprintf("\nexport AGENT_ROOT=%s\n", root)
	f, err := os.OpenFile(cfg, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("opening %s: %w", cfg, err)
	}
	defer f.Close()
	if _, err := f.WriteString(line); err != nil {
		return fmt.Errorf("writing to %s: %w", cfg, err)
	}
	return nil
}
