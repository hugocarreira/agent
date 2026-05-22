package config

import (
	"os"
	"path/filepath"
)

func AgentRoot() string {
	if env := os.Getenv("AGENT_ROOT"); env != "" {
		return env
	}
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		marker := filepath.Join(dir, "cli")
		if fi, err := os.Stat(marker); err == nil && fi.IsDir() {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

func ProjectDir(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	wd, err := os.Getwd()
	if err != nil {
		return path
	}
	return filepath.Join(wd, path)
}
