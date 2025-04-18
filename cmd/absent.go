
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var absentCommand = &cobra.Command{
	Use: "absent",
	Short: "absent a course",
	Long: "longer text",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("hello from absent")
	},
}
