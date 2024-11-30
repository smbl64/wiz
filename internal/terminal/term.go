package terminal

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func StdinHasData() (bool, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false, fmt.Errorf("failed to stat stdin: %v", err)
	}

	isPipe := fi.Mode()&os.ModeNamedPipe == os.ModeNamedPipe
	return isPipe, nil
}

func ReadStdinIfData() (string, error) {
	hasData, err := StdinHasData()
	if err != nil {
		return "", err
	}

	if !hasData {
		return "", nil
	}

	si, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", fmt.Errorf("failed to read stdin: %v", err)
	}

	return strings.TrimSuffix(string(si), "\n"), nil
}
