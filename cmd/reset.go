package cmd

import (
	"github.com/hababisha/ejigo/models"
	"github.com/hababisha/ejigo/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset all courses",
	Run: func(cmd *cobra.Command, args []string) {
		utils.SaveCourses(make(map[string]*models.Course))
		fmt.Println("All courses reset")
	},
}
