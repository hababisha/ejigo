package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:"Ejigo",
		Short: "Absent tracker",
		Long: `Skip classes with freedom`,	
	}
)

func init(){
	rootCmd.AddCommand(resetCommand)
	rootCmd.AddCommand(absentCommand)
	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(summaryCommand)
}

func Execute() error{
	return rootCmd.Execute()
}