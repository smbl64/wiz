package main

import (
	"fmt"
	"os"

	"github.com/smbl64/wiz/cmd"
	"github.com/smbl64/wiz/internal/config"
)

func main() {
	if err := config.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	cmd.Execute()
}
