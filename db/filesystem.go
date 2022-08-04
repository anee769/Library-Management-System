package db

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const dbFolder = "data"

func Exists() bool {
	manifest := filepath.Join(Dir(), "MANIFEST")

	if _, err := os.Stat(manifest); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func Dir() string {
	executable, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("database directory detection failure: exec path detection failure: %w", err))
	}

	execDir := filepath.Dir(executable)
	return filepath.Join(execDir, dbFolder)
}
