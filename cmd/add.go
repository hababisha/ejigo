package cmd

import (
	"fmt"
	"github.com/hababisha/ejigo/models"
	"github.com/hababisha/ejigo/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [course name]",
	Short: "Add a new course",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		courses, _ := utils.GetCourses()
		if _, exists := courses[name]; exists {
			fmt.Println("course already exists")
			return
		}

		courses[name] = &models.Course{Name: name, Absences: 0}
		utils.SaveCourses(courses)
		fmt.Printf("added course: %s\n", name)
	},
}
