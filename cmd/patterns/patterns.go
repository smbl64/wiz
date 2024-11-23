package patterns

import "github.com/spf13/cobra"

func Initialize(rootCmd *cobra.Command) {
	var patternsCmd = &cobra.Command{
		Use:   "patterns",
		Short: "Manage patterns",
	}
	rootCmd.AddCommand(patternsCmd)

	patternsCmd.AddCommand(newCmd)
	patternsCmd.AddCommand(listCmd)
	patternsCmd.AddCommand(editCmd)
	patternsCmd.AddCommand(delCmd)

}
