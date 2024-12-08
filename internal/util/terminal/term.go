package terminal

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

func Confirm(question string) bool {
	r := bufio.NewReader(os.Stdin)
	tries := 10

	for ; tries > 0; tries-- {
		fmt.Printf("%s [y/n]: ", question)

		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Empty input (i.e. "\n")
		if len(res) < 2 {
			continue
		}

		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}

	return false
}
