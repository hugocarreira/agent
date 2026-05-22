package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAgentRoot_default(t *testing.T) {
	os.Unsetenv("AGENT_ROOT")
	td := t.TempDir()
	os.MkdirAll(filepath.Join(td, "cli"), 0755)

	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	got := AgentRoot()
	if got != td {
		t.Errorf("AgentRoot() = %q, want %q", got, td)
	}
}

func TestAgentRoot_env(t *testing.T) {
	os.Setenv("AGENT_ROOT", "/custom/path")
	defer os.Unsetenv("AGENT_ROOT")
	got := AgentRoot()
	if got != "/custom/path" {
		t.Errorf("AgentRoot() = %q, want %q", got, "/custom/path")
	}
}

func TestAgentRoot_notFound(t *testing.T) {
	os.Unsetenv("AGENT_ROOT")
	td := t.TempDir()

	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	got := AgentRoot()
	if got != "" {
		t.Errorf("AgentRoot() = %q, want \"\"", got)
	}
}

func TestProjectDir_relative(t *testing.T) {
	td := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	want := filepath.Join(td, "myproject")
	got := ProjectDir("myproject")
	if got != want {
		t.Errorf("ProjectDir() = %q, want %q", got, want)
	}
}

func TestProjectDir_absolute(t *testing.T) {
	got := ProjectDir("/absolute/path")
	if got != "/absolute/path" {
		t.Errorf("ProjectDir() = %q, want %q", got, "/absolute/path")
	}
}
