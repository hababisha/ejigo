
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var summaryCommand = &cobra.Command{
	Use: "summary",
	Short: "display summary",
	Long: "longer text",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("hello from summary")
	},
}
