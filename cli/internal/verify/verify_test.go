package verify

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckSymlink_notFound(t *testing.T) {
	td := t.TempDir()
	ok := true
	checkSymlink(filepath.Join(td, "nonexistent"), &ok)
	if ok {
		t.Error("ok should be false when file not found")
	}
}

func TestCheckSymlink_notASymlink(t *testing.T) {
	td := t.TempDir()
	file := filepath.Join(td, "regular.txt")
	os.WriteFile(file, []byte("hello"), 0644)
	ok := true
	checkSymlink(file, &ok)
	if ok {
		t.Error("ok should be false for regular file")
	}
}

func TestCheckSymlink_validSymlink(t *testing.T) {
	td := t.TempDir()
	target := filepath.Join(td, "target.txt")
	os.WriteFile(target, []byte("hello"), 0644)
	link := filepath.Join(td, "link.txt")
	os.Symlink(target, link)
	ok := true
	checkSymlink(link, &ok)
	if !ok {
		t.Error("ok should be true for valid symlink")
	}
}

func TestCheckSymlink_brokenSymlink(t *testing.T) {
	td := t.TempDir()
	link := filepath.Join(td, "broken.txt")
	os.Symlink("/nonexistent/target", link)
	ok := true
	checkSymlink(link, &ok)
	// Broken symlink still has ModeSymlink set
	if !ok {
		t.Error("ok should be true for symlink (even broken)")
	}
}

func TestProjectPackEnabled(t *testing.T) {
	td := t.TempDir()
	if projectPackEnabled(td) {
		t.Fatal("projectPackEnabled should be false without state files")
	}

	os.WriteFile(filepath.Join(td, "FEATURES.yaml"), []byte("features: []\n"), 0644)
	if !projectPackEnabled(td) {
		t.Fatal("projectPackEnabled should be true when FEATURES.yaml exists")
	}
}

func TestValidateProjectPack_valid(t *testing.T) {
	td := t.TempDir()
	os.WriteFile(filepath.Join(td, "AGENTS.md"), []byte("# Agent Instructions\n"), 0644)
	os.WriteFile(filepath.Join(td, "PROGRESS.md"), []byte("# Progress\n"), 0644)
	os.WriteFile(filepath.Join(td, "FEATURES.yaml"), []byte(validFeaturesYAML()), 0644)

	errors := validateProjectPack(td)
	if len(errors) != 0 {
		t.Fatalf("validateProjectPack returned errors: %v", errors)
	}
}

func TestValidateProjectPack_missingFiles(t *testing.T) {
	td := t.TempDir()
	os.WriteFile(filepath.Join(td, "FEATURES.yaml"), []byte("features: []\n"), 0644)

	errors := validateProjectPack(td)
	if len(errors) != 2 {
		t.Fatalf("expected 2 missing file errors, got %d: %v", len(errors), errors)
	}
	if !strings.Contains(errors[0], "AGENTS.md") && !strings.Contains(errors[1], "AGENTS.md") {
		t.Fatal("missing AGENTS.md error not reported")
	}
	if !strings.Contains(errors[0], "PROGRESS.md") && !strings.Contains(errors[1], "PROGRESS.md") {
		t.Fatal("missing PROGRESS.md error not reported")
	}
}

func TestValidateFeaturesFile_invalidEntries(t *testing.T) {
	td := t.TempDir()
	featuresFile := filepath.Join(td, "FEATURES.yaml")
	os.WriteFile(featuresFile, []byte(`features:
  - id: ""
    behavior: ""
    done_when: []
    state: partial
`), 0644)

	errors := validateFeaturesFile(featuresFile)
	if len(errors) != 4 {
		t.Fatalf("expected 4 errors, got %d: %v", len(errors), errors)
	}
}

func TestValidateFeaturesFile_invalidYAML(t *testing.T) {
	td := t.TempDir()
	featuresFile := filepath.Join(td, "FEATURES.yaml")
	os.WriteFile(featuresFile, []byte("features:\n  - id: [\n"), 0644)

	errors := validateFeaturesFile(featuresFile)
	if len(errors) != 1 || !strings.Contains(errors[0], "invalid YAML") {
		t.Fatalf("expected invalid YAML error, got %v", errors)
	}
}

func validFeaturesYAML() string {
	return `features:
  - id: pricing-page
    behavior: Marketing site has a pricing page
    done_when:
      - /pricing route renders
      - CTA links to signup
    check:
      type: manual
      note: Run app and verify in browser.
    state: active
`
}
