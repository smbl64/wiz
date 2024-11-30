package patterns

import (
	"os"
	"os/exec"

	"github.com/smbl64/wiz/internal/patmgr"
	"github.com/spf13/cobra"
)

const fallbackEditor = "nano"

var editCmd = &cobra.Command{
	Use:                   "edit <pattern>",
	Short:                 "Edit a pattern",
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     patternFlagCompletion,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return cmd.Usage()
		}

		editor := getDefaultEditor()

		systemFile, err := patmgr.Default().GetSystemFileName(args[0])
		if err != nil {
			cmd.PrintErr(err)
			return nil
		}

		c := exec.Command(editor, systemFile)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()

		return nil
	},
}

func getDefaultEditor() string {
	envVars := []string{"VISUAL", "EDITOR"}

	for _, ev := range envVars {
		editor, found := os.LookupEnv(ev)
		if editor != "" && found {
			return editor
		}
	}

	return fallbackEditor
}
