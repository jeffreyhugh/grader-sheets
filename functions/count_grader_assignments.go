package functions

import (
	"../structures"
)

func CountGraderAssignments(graders *[]structures.Grader) int {
	graderAssignmentsCount := 0

	for _, g := range *graders {
		if g.Grade == 0 {
			graderAssignmentsCount += structures.DefaultGrade
		} else {
			graderAssignmentsCount += g.Grade
		}
	}

	return graderAssignmentsCount
}
