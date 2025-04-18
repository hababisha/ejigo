
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var addCommand = &cobra.Command{
	Use: "add",
	Short: "add course",
	Long: "longer text",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("hello from add")
	},
}
