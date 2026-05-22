package harness

import (
	"os"
	"os/exec"
	"path/filepath"
)

type Agent struct {
	Name       string
	Dir        string
	ConfigFile string
	Installed  bool
}

func AllAgents() []Agent {
	home := os.Getenv("HOME")
	all := []Agent{
		{Name: "OpenCode", Dir: filepath.Join(home, ".config", "opencode"), ConfigFile: "AGENTS.md"},
		{Name: "Codex", Dir: filepath.Join(home, ".codex"), ConfigFile: "AGENTS.md"},
		{Name: "Claude", Dir: filepath.Join(home, ".claude"), ConfigFile: "CLAUDE.md"},
		{Name: "GitHub Copilot", Dir: filepath.Join(home, ".copilot"), ConfigFile: "copilot-instructions.md"},
		{Name: "Hermes", Dir: filepath.Join(home, ".hermes"), ConfigFile: "SOUL.md"},
		{Name: "Gemini CLI", Dir: filepath.Join(home, ".gemini"), ConfigFile: "GEMINI.md"},
	}
	for i := range all {
		if _, err := os.Stat(all[i].Dir); err == nil {
			all[i].Installed = true
		}
	}
	return all
}

func DetectAgents() []Agent {
	all := AllAgents()
	var installed []Agent
	for _, a := range all {
		if a.Installed {
			installed = append(installed, a)
		}
	}
	return installed
}

func RTKVersion() string {
	out, err := exec.Command("rtk", "--version").Output()
	if err != nil {
		return ""
	}
	return string(out)
}
