package cmd

import (
	"fmt"

	"github.com/hababisha/ejigo/utils"
	"github.com/spf13/cobra"
)

var absentCmd = &cobra.Command{
	Use:   "absent [course name]",
	Short: "increment absence count for a course",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		courses, _ := utils.GetCourses()
		if course, exists := courses[name]; exists {
			course.Absences += absentCount
			utils.SaveCourses(courses)
			fmt.Printf("You got absent for course: %s\n", name)
			fmt.Printf("Total: %d", course.Absences)
		} else {
			fmt.Println("Course not found")
		}
	},
}
