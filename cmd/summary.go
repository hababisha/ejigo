package cmd

import (
	"fmt"
	"github.com/hababisha/ejigo/utils"
	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Display all your courses and absence counts",
	Run: func(cmd *cobra.Command, args []string) {
		courses, _ := utils.GetCourses()
		fmt.Println("summary")
		fmt.Println("--------")
		if len(courses) == 0 {
			fmt.Println("No courses found")
			return
		}

		for _, c := range courses {
			fmt.Printf("%s: %d absences\n", c.Name, c.Absences)
		}
	},
}

