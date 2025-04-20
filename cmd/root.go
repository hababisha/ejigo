package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "ejigo",
		Short: "Absent tracker",
		Long:  `Skip classes with freedom`,
	}

	defaultAbsentCount int
	absentCount        int
)

func init() {
	addCmd.PersistentFlags().IntVar(&defaultAbsentCount, "absent", 1, "default absent out for course")
	absentCmd.PersistentFlags().IntVar(&absentCount, "count", 1, "absent count")

	rootCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(absentCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(summaryCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() error {
	return rootCmd.Execute()
}
