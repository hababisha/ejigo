
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var resetCommand = &cobra.Command{
	Use: "reset",
	Short: "reset all courses",
	Long: "longer text",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("hello from reset")
	},
}
