package verify

import (
	"os"
	"path/filepath"
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


