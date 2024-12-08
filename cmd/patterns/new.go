package patterns

import (
	"os"
	"unicode/utf8"

	"github.com/smbl64/wiz/internal/patmgr"
	"github.com/smbl64/wiz/internal/util/terminal"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:                   "new <pattern_name>",
	Short:                 "Create a new pattern",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return cmd.Usage()
		}

		patName := args[0]
		mgr := patmgr.Default()

		exists, err := mgr.Exists(patName)
		if err != nil {
			cmd.PrintErr(err)
			return nil
		}

		if exists {
			edit := terminal.Confirm("A pattern with this name already exists. Open to edit?")
			if edit {
				return openPatternForEdit(cmd, patName)
			}

			return nil
		}

		return createNewPattern(cmd, patName)
	},
}

func createNewPattern(cmd *cobra.Command, patName string) error {
	tempFile, err := makeTempPatternFile(patName)
	if err != nil {
		cmd.PrintErr(err)
		return nil
	}

	if err = editFile(tempFile); err != nil {
		cmd.PrintErr(err)
		return nil
	}

	// If file has any content, ask pattern manager to create a new one
	bytes, err := os.ReadFile(tempFile)
	if err != nil {
		cmd.PrintErr(err)
		return nil
	}

	if utf8.RuneCount(bytes) == 0 {
		cmd.Println("Pattern is empty. Will skip creating a new pattern.")
		return nil
	}

	err = patmgr.Default().Create(bytes, patName)
	if err != nil {
		cmd.PrintErr(err)
		return nil
	}

	cmd.Printf("Added '%s' to the pattern repository.", patName)

	return nil
}

func makeTempPatternFile(patName string) (string, error) {
	fp, err := os.CreateTemp("", patName+"-*.md")
	if err != nil {
		return "", err
	}

	fp.Close()
	return fp.Name(), nil
}
