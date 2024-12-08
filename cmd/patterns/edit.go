package patterns

import (
	"os"
	"os/exec"

	"github.com/smbl64/wiz/internal/patmgr"
	"github.com/smbl64/wiz/internal/util/flags"
	"github.com/spf13/cobra"
)

const fallbackEditor = "nano"

var editCmd = &cobra.Command{
	Use:                   "edit <pattern>",
	Short:                 "Edit a pattern",
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     flags.PatternsFlagCompletionFunc,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return cmd.Usage()
		}

		patName := args[0]
		return openPatternForEdit(cmd, patName)
	},
}

func openPatternForEdit(cmd *cobra.Command, patterName string) error {
	systemFile, err := patmgr.Default().GetSystemFileName(patterName)
	if err != nil {
		cmd.PrintErr(err)
		return nil
	}

	if err = editFile(systemFile); err != nil {
		cmd.PrintErr(cmd)
	}

	return nil
}

func editFile(filePath string) error {
	editor := getDefaultEditor()
	c := exec.Command(editor, filePath)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()

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
