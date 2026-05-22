package harness

import (
	"os"
	"path/filepath"

	"github.com/hugocarreira/agentrc/cli/internal/log"
)

func Setup(root string) {
	log.L.Info("🔗 Setting up AI coding agents...")
	log.L.Info("Philosophy: Just talk to it. No orchestrators, no charades.")

	rtk := RTKVersion()
	if rtk != "" {
		log.L.Info("✅ RTK: " + rtk)
	} else {
		log.L.Info("ℹ️  RTK not found — optional, install with: brew install rtk-ai/tap/rtk")
	}

	agents := DetectAgents()
	if len(agents) == 0 {
		log.L.Info("No AI coding agents detected. Install one of:")
		log.L.Info("  • OpenCode:       https://opencode.ai")
		log.L.Info("  • Codex:          https://chatgpt.com/codex/")
		log.L.Info("  • Claude:         https://claude.ai")
		log.L.Info("  • GitHub Copilot: https://github.com/features/copilot")
		log.L.Info("  • Hermes:         https://hermes-agent.nousresearch.com")
		log.L.Info("  • Gemini CLI:     https://cloud.google.com/gemini-cli")
		return
	}

	content := "@" + root + "/RTK.md\n\n@" + root + "/AGENTS.md\n"

	for _, a := range agents {
		cfgPath := filepath.Join(a.Dir, a.ConfigFile)
		os.Remove(cfgPath)
		if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
			log.L.Info("  ❌ " + a.Name + " config: " + err.Error())
			continue
		}
		log.L.Info("  ↳ " + a.Name + " → " + cfgPath)

		skillsDir := filepath.Join(a.Dir, "skills")
		os.RemoveAll(skillsDir)
		if err := os.Symlink(filepath.Join(root, "skills"), skillsDir); err != nil {
			log.L.Info("    ⚠️  skills/: " + err.Error())
		} else {
			log.L.Info("    ↳ skills/ → " + root + "/skills")
		}

		if a.Name == "Claude" {
			agentsLink := filepath.Join(a.Dir, "AGENTS.md")
			os.Remove(agentsLink)
			os.Symlink(filepath.Join(a.Dir, "CLAUDE.md"), agentsLink)
			log.L.Info("    ↳ AGENTS.md → CLAUDE.md")
		}
	}

	opencodeDir := filepath.Join(os.Getenv("HOME"), ".config", "opencode")
	if _, err := os.Stat(opencodeDir); err == nil {
		pluginsDir := filepath.Join(opencodeDir, "plugins")
		os.RemoveAll(pluginsDir)
		os.Symlink(filepath.Join(root, "plugins"), pluginsDir)
		log.L.Info("  ↳ plugins/ → linked")
	}

	oldAgents := filepath.Join(os.Getenv("HOME"), ".config", "opencode", "agents")
	if fi, err := os.Lstat(oldAgents); err == nil && fi.Mode()&os.ModeSymlink != 0 {
		os.Remove(oldAgents)
		log.L.Info("  ↳ Removed old orchestrator agents/")
	}

	log.L.Info("✅ Config + skills setup complete!")
}