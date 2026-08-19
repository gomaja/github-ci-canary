package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("writes generated model", func(t *testing.T) {
		workDir := t.TempDir()
		if err := os.Mkdir(filepath.Join(workDir, "generated"), 0o700); err != nil {
			t.Fatalf("create generated directory: %v", err)
		}
		t.Chdir(workDir)

		if code := run(); code != 0 {
			t.Fatalf("run() = %d, want 0", code)
		}

		got, err := os.ReadFile(filepath.Join(workDir, "generated", "model.go"))
		if err != nil {
			t.Fatalf("read generated model: %v", err)
		}
		if string(got) != generatedModel {
			t.Fatalf("generated model = %q, want %q", got, generatedModel)
		}
	})

	t.Run("reports write failure", func(t *testing.T) {
		t.Chdir(t.TempDir())

		if code := run(); code != 1 {
			t.Fatalf("run() = %d, want 1", code)
		}
	})
}
