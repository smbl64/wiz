// patmgr is the package for loading and managing patterns.
package patmgr

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/smbl64/wiz/internal/config"
	"github.com/smbl64/wiz/internal/util/paths"
)

var ErrNotExist = errors.New("Pattern does not exist")

type PatternManager struct {
	rootDir string
}

// New creates a new PatternManager
func New(rootDir string) *PatternManager {
	return &PatternManager{rootDir: rootDir}
}

func Default() *PatternManager {
	return &PatternManager{
		rootDir: path.Join(config.ConfigDir(), "patterns"),
	}
}

func (m *PatternManager) List() ([]string, error) {

	entries, err := os.ReadDir(m.rootDir)
	if err != nil {
		return nil, fmt.Errorf("list patterns: %v", err)
	}

	result := make([]string, len(entries))
	for i, e := range entries {
		if strings.HasPrefix(e.Name(), ".") {
			continue
		}

		if !e.IsDir() {
			continue
		}

		result[i] = e.Name()
	}

	return result, nil
}

func (m *PatternManager) Load(patternName string) (string, error) {
	fullPath, err := m.GetSystemFileName(patternName)
	if err != nil {
		return "", err
	}

	bb, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}

	return string(bb), nil
}

func (m *PatternManager) GetSystemFileName(patternName string) (string, error) {
	fullPath := filepath.Join(m.rootDir, patternName, "system.md")
	exist, err := paths.Exists(fullPath)
	if err != nil {
		return "", err
	}

	if !exist {
		return "", ErrNotExist
	}

	return fullPath, nil
}
