package verify

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hugocarreira/agentrc/cli/internal/harness"
	"github.com/hugocarreira/agentrc/cli/internal/log"
	"gopkg.in/yaml.v3"
)

var allowedFeatureStates = map[string]struct{}{
	"not_started": {},
	"active":      {},
	"blocked":     {},
	"done":        {},
}

type featureList struct {
	Features []feature `yaml:"features"`
}

type feature struct {
	ID       string        `yaml:"id"`
	Behavior string        `yaml:"behavior"`
	DoneWhen []string      `yaml:"done_when"`
	Check    *featureCheck `yaml:"check,omitempty"`
	State    string        `yaml:"state"`
}

type featureCheck struct {
	Type string `yaml:"type,omitempty"`
	Note string `yaml:"note,omitempty"`
}

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

	projectDir, err := os.Getwd()
	if err == nil {
		log.L.Info("")
		log.L.Info("--- Current Project ---")
		if projectPackEnabled(projectDir) {
			errors := validateProjectPack(projectDir)
			if len(errors) == 0 {
				log.L.Info("  ✅ " + projectDir + " project pack")
			} else {
				allOk = false
				for _, msg := range errors {
					log.L.Info("  ❌ " + msg)
				}
			}
		} else {
			log.L.Info("  ℹ️  no FEATURES.yaml or PROGRESS.md in current directory — skipping project checks")
		}
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

func projectPackEnabled(projectDir string) bool {
	return fileExists(filepath.Join(projectDir, "FEATURES.yaml")) || fileExists(filepath.Join(projectDir, "PROGRESS.md"))
}

func validateProjectPack(projectDir string) []string {
	var errors []string

	agentsPath := filepath.Join(projectDir, "AGENTS.md")
	featuresPath := filepath.Join(projectDir, "FEATURES.yaml")
	progressPath := filepath.Join(projectDir, "PROGRESS.md")

	if !fileExists(agentsPath) {
		errors = append(errors, missingFileMessage(agentsPath))
	}
	if !fileExists(featuresPath) {
		errors = append(errors, missingFileMessage(featuresPath))
	}
	if !fileExists(progressPath) {
		errors = append(errors, missingFileMessage(progressPath))
	}
	if !fileExists(featuresPath) {
		return errors
	}

	return append(errors, validateFeaturesFile(featuresPath)...)
}

func validateFeaturesFile(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		return []string{fmt.Sprintf("%s unreadable: %v", path, err)}
	}

	var features featureList
	if err := yaml.Unmarshal(data, &features); err != nil {
		return []string{fmt.Sprintf("%s invalid YAML: %v", path, err)}
	}

	var errors []string
	for i, feature := range features.Features {
		label := fmt.Sprintf("%s feature[%d]", path, i)

		if strings.TrimSpace(feature.ID) == "" {
			errors = append(errors, label+" missing id")
		}
		if strings.TrimSpace(feature.Behavior) == "" {
			errors = append(errors, label+" missing behavior")
		}
		if len(nonEmptyLines(feature.DoneWhen)) == 0 {
			errors = append(errors, label+" missing done_when")
		}
		if _, ok := allowedFeatureStates[feature.State]; !ok {
			errors = append(errors, label+" has invalid state "+feature.State)
		}
	}

	return errors
}

func nonEmptyLines(lines []string) []string {
	var filtered []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

func missingFileMessage(path string) string {
	return "missing file: " + path
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
