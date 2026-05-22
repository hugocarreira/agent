package harness

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAllAgents_returnsSix(t *testing.T) {
	agents := AllAgents()
	if len(agents) != 6 {
		t.Errorf("AllAgents() returned %d, want 6", len(agents))
	}
}

func TestAllAgents_fields(t *testing.T) {
	home := os.Getenv("HOME")
	agents := AllAgents()

	tests := []struct {
		name       string
		wantDir    string
		wantConfig string
	}{
		{"OpenCode", filepath.Join(home, ".config", "opencode"), "AGENTS.md"},
		{"Codex", filepath.Join(home, ".codex"), "AGENTS.md"},
		{"Claude", filepath.Join(home, ".claude"), "CLAUDE.md"},
		{"GitHub Copilot", filepath.Join(home, ".copilot"), "copilot-instructions.md"},
		{"Hermes", filepath.Join(home, ".hermes"), "SOUL.md"},
		{"Gemini CLI", filepath.Join(home, ".gemini"), "GEMINI.md"},
	}
	for _, tt := range tests {
		var found *Agent
		for _, a := range agents {
			if a.Name == tt.name {
				found = &a
				break
			}
		}
		if found == nil {
			t.Errorf("agent %q not found", tt.name)
			continue
		}
		if found.Dir != tt.wantDir {
			t.Errorf("%s Dir = %q, want %q", tt.name, found.Dir, tt.wantDir)
		}
		if found.ConfigFile != tt.wantConfig {
			t.Errorf("%s ConfigFile = %q, want %q", tt.name, found.ConfigFile, tt.wantConfig)
		}
	}
}

func TestDetectAgents_onlyInstalled(t *testing.T) {
	// With no agents installed, should return empty
	home := os.Getenv("HOME")
	defer os.Setenv("HOME", home)

	td := t.TempDir()
	os.Setenv("HOME", td)

	agents := DetectAgents()
	for _, a := range agents {
		t.Errorf("DetectAgents() returned %q, but agent dir should not exist", a.Name)
	}
}

func TestDetectAgents_oneInstalled(t *testing.T) {
	td := t.TempDir()
	home := os.Getenv("HOME")
	defer os.Setenv("HOME", home)
	os.Setenv("HOME", td)

	os.MkdirAll(filepath.Join(td, ".codex"), 0755)

	agents := DetectAgents()
	if len(agents) != 1 {
		t.Fatalf("DetectAgents() returned %d, want 1", len(agents))
	}
	if agents[0].Name != "Codex" {
		t.Errorf("expected Codex, got %s", agents[0].Name)
	}
}

func TestAllAgents_installedFlag(t *testing.T) {
	td := t.TempDir()
	home := os.Getenv("HOME")
	defer os.Setenv("HOME", home)
	os.Setenv("HOME", td)

	os.MkdirAll(filepath.Join(td, ".claude"), 0755)

	agents := AllAgents()
	for _, a := range agents {
		switch a.Name {
		case "Claude":
			if !a.Installed {
				t.Errorf("Claude should be installed")
			}
		default:
			if a.Installed {
				t.Errorf("%s should not be installed", a.Name)
			}
		}
	}
}
