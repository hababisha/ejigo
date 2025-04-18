package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "ejigo",
		Short: "Absent tracker",
		Long:  `Skip classes with freedom`,
	}
)

func init() {
	rootCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(absentCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(summaryCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() error {
	return rootCmd.Execute()
}
