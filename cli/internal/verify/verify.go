package verify

import (
	"os"
	"path/filepath"

	"github.com/hugocarreira/agentrc/cli/internal/harness"
	"github.com/hugocarreira/agentrc/cli/internal/log"
)

func Run() {
	log.L.Info("🔍 Verification Report")
	log.L.Info("")

	rtk := harness.RTKVersion()
	if rtk != "" {
		log.L.Info("✅ RTK: " + rtk)
	} else {
		log.L.Info("ℹ️  RTK not found — optional")
	}
	log.L.Info("")

	allOk := true
	agents := harness.AllAgents()

	for _, a := range agents {
		log.L.Info("")
		log.L.Info("--- " + a.Name + " ---")

		if !a.Installed {
			log.L.Info("  ⏭️  not installed")
			continue
		}

		cfgPath := filepath.Join(a.Dir, a.ConfigFile)
		checkFile(cfgPath, &allOk)

		skillsPath := filepath.Join(a.Dir, "skills")
		checkSymlink(skillsPath, &allOk)

		if a.Name == "Claude" {
			agentsLink := filepath.Join(a.Dir, "AGENTS.md")
			checkSymlink(agentsLink, &allOk)
		}
	}

	opencodeDir := filepath.Join(os.Getenv("HOME"), ".config", "opencode")
	if _, err := os.Stat(opencodeDir); err == nil {
		log.L.Info("")
		log.L.Info("--- OpenCode Plugins ---")
		checkSymlink(filepath.Join(opencodeDir, "plugins"), &allOk)
	}

	log.L.Info("")
	if allOk {
		log.L.Info("✅ All checks passed! Ready to go.")
	} else {
		log.L.Info("⚠️  Some checks failed. Run: agentrc setup")
	}
}

func checkFile(path string, ok *bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.L.Info("  ❌ not found: " + path)
		*ok = false
		return
	}
	log.L.Info("  ✅ " + path)
}

func checkSymlink(path string, ok *bool) {
	fi, err := os.Lstat(path)
	if os.IsNotExist(err) {
		log.L.Info("  ❌ not found: " + path)
		*ok = false
		return
	}
	if fi.Mode()&os.ModeSymlink == 0 {
		log.L.Info("  ⚠️  not a symlink: " + path)
		*ok = false
		return
	}
	target, _ := os.Readlink(path)
	log.L.Info("  ✅ " + path + " → " + target)
}
