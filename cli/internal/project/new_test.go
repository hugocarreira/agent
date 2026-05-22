package project

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNew_createsAgentsMd(t *testing.T) {
	td := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	projectDir := filepath.Join(td, "testproj")
	os.MkdirAll(projectDir, 0755)

	root := filepath.Join(td, "agent")
	os.MkdirAll(root, 0755)
	tmpl := filepath.Join(root, "AGENTS.template.md")
	os.WriteFile(tmpl, []byte("TEMPLATE CONTENT\n"), 0644)

	New(root, "testproj")

	agentsFile := filepath.Join(projectDir, "AGENTS.md")
	data, err := os.ReadFile(agentsFile)
	if err != nil {
		t.Fatalf("AGENTS.md not created: %v", err)
	}
	if string(data) != "READ "+root+"/AGENTS.md BEFORE ANYTHING.\n\nTEMPLATE CONTENT\n" {
		t.Errorf("unexpected content:\n%q", string(data))
	}
}

func TestNew_projectNotExists(t *testing.T) {
	td := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	root := filepath.Join(td, "agent")
	os.MkdirAll(root, 0755)

	New(root, "nonexistent")

	agentsFile := filepath.Join(td, "nonexistent", "AGENTS.md")
	if _, err := os.Stat(agentsFile); !os.IsNotExist(err) {
		t.Error("AGENTS.md should not exist for missing project")
	}
}

func TestNew_missingTemplate(t *testing.T) {
	td := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	projectDir := filepath.Join(td, "testproj")
	os.MkdirAll(projectDir, 0755)

	New(filepath.Join(td, "agent"), "testproj")

	agentsFile := filepath.Join(projectDir, "AGENTS.md")
	if _, err := os.Stat(agentsFile); !os.IsNotExist(err) {
		t.Error("AGENTS.md should not exist when template missing")
	}
}

func TestNew_backupExisting(t *testing.T) {
	td := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	projectDir := filepath.Join(td, "testproj")
	os.MkdirAll(projectDir, 0755)

	root := filepath.Join(td, "agent")
	os.MkdirAll(root, 0755)
	tmpl := filepath.Join(root, "AGENTS.template.md")
	os.WriteFile(tmpl, []byte("NEW CONTENT\n"), 0644)

	existing := filepath.Join(projectDir, "AGENTS.md")
	os.WriteFile(existing, []byte("ORIGINAL\n"), 0644)

	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	w.WriteString("y\n")
	w.Close()
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	New(root, "testproj")
	os.Stdin = oldStdin

	data, _ := os.ReadFile(existing)
	expected := "READ " + root + "/AGENTS.md BEFORE ANYTHING.\n\nNEW CONTENT\n"
	if strings.TrimSpace(string(data)) != strings.TrimSpace(expected) {
		t.Errorf("existing content = %q, want %q", strings.TrimSpace(string(data)), strings.TrimSpace(expected))
	}

	files, _ := filepath.Glob(filepath.Join(projectDir, "AGENTS.md.backup.*"))
	if len(files) == 0 {
		t.Error("backup file not created")
	}
}
