// patmgr is the package for loading and managing patterns.
package patmgr

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/smbl64/wiz/internal/config"
)

type PatternManager struct {
	rootDir string
}

// New creates a new PatternManager
func New(rootDir string) *PatternManager {
	return &PatternManager{rootDir: rootDir}
}

func Default() *PatternManager {
	return &PatternManager{
		rootDir: config.Current().PatternsDir,
	}
}

func (m *PatternManager) List() ([]string, error) {

	entries, err := os.ReadDir(m.rootDir)
	if err != nil {
		return nil, fmt.Errorf("list patterns: %v", err)
	}

	result := make([]string, len(entries))
	for i, e := range entries {
		result[i] = e.Name()
	}

	return result, nil
}

func (m *PatternManager) Load(patternName string) (string, error) {
	fullPath := filepath.Join(m.rootDir, patternName, "system.md")

	bb, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}

	return string(bb), nil
}
