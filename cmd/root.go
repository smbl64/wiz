package cmd

import (
	"fmt"
	"os"

	"github.com/smbl64/wiz/cmd/config"
	"github.com/smbl64/wiz/cmd/generate"
	"github.com/smbl64/wiz/cmd/patterns"
	"github.com/smbl64/wiz/cmd/tools"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wiz",
	Short: "wiz is a little helper to work with LLMs",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	config.Initialize(rootCmd)
	generate.Initialize(rootCmd)
	patterns.Initialize(rootCmd)
	tools.Initialize(rootCmd)
}
