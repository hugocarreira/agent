package project

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLink_addsReference(t *testing.T) {
	td := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	os.Setenv("AGENT_ROOT", td)
	defer os.Unsetenv("AGENT_ROOT")

	projectDir := filepath.Join(td, "testproj")
	os.MkdirAll(projectDir, 0755)
	existing := filepath.Join(projectDir, "AGENTS.md")
	os.WriteFile(existing, []byte("# My Project\n"), 0644)

	Link("testproj")

	data, _ := os.ReadFile(existing)
	content := string(data)
	if !strings.Contains(content, "READ "+td+"/AGENTS.md") {
		t.Errorf("expected global ref in content, got:\n%s", content)
	}
	if !strings.Contains(content, "# My Project") {
		t.Errorf("expected original content preserved, got:\n%s", content)
	}
}

func TestLink_noAgentsMd(t *testing.T) {
	td := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	projectDir := filepath.Join(td, "testproj")
	os.MkdirAll(projectDir, 0755)

	Link("testproj")

	if _, err := os.Stat(filepath.Join(projectDir, "AGENTS.md")); !os.IsNotExist(err) {
		t.Error("AGENTS.md should not be created when it doesn't exist")
	}
}

func TestLink_alreadyReferenced(t *testing.T) {
	td := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	os.Setenv("AGENT_ROOT", td)
	defer os.Unsetenv("AGENT_ROOT")

	projectDir := filepath.Join(td, "testproj")
	os.MkdirAll(projectDir, 0755)
	existing := filepath.Join(projectDir, "AGENTS.md")
	content := "READ " + td + "/AGENTS.md BEFORE ANYTHING.\n\n# My Project\n"
	os.WriteFile(existing, []byte(content), 0644)

	Link("testproj")

	data, _ := os.ReadFile(existing)
	if string(data) != content {
		t.Errorf("content changed for already-referenced file:\n%q", string(data))
	}
}

func TestLink_backupCreated(t *testing.T) {
	td := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	os.Setenv("AGENT_ROOT", td)
	defer os.Unsetenv("AGENT_ROOT")

	projectDir := filepath.Join(td, "testproj")
	os.MkdirAll(projectDir, 0755)
	existing := filepath.Join(projectDir, "AGENTS.md")
	os.WriteFile(existing, []byte("# Original\n"), 0644)

	Link("testproj")

	files, _ := filepath.Glob(filepath.Join(projectDir, "AGENTS.md.backup.*"))
	if len(files) == 0 {
		t.Error("backup file not created")
	}
}

func TestLink_noRoot(t *testing.T) {
	os.Unsetenv("AGENT_ROOT")
	td := t.TempDir()
	origWd, _ := os.Getwd()
	os.Chdir(td)
	defer os.Chdir(origWd)

	projectDir := filepath.Join(td, "testproj")
	os.MkdirAll(projectDir, 0755)
	existing := filepath.Join(projectDir, "AGENTS.md")
	os.WriteFile(existing, []byte("# My Project\n"), 0644)

	Link("testproj")

	data, _ := os.ReadFile(existing)
	if string(data) != "# My Project\n" {
		t.Error("file should be unchanged when root not found")
	}
}
