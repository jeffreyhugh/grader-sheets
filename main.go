package main

import (
	"./functions"
	"fmt"
)

func main() {
	graders, err := functions.UnmarshalGraders("") // TODO command-line argument to pass graderJson
	if err != nil {
		panic(err)
	}

	graderAssignmentCount := functions.CountGraderAssignments(graders.G)
	if graderAssignmentCount == 0 {
		panic("0 grader assignments counted")
	}

	submissionCount, err := functions.CountSubmissions()
	if err != nil {
		panic(err)
	} else if submissionCount == 0 {
		panic("0 submissions detected")
	}

	if submissionCount == graderAssignmentCount {
		fmt.Printf("Grader Assignment Count == Submission Count (%d == %d), continuing\n", graderAssignmentCount, submissionCount)
	} else if submissionCount > graderAssignmentCount {
		fmt.Printf("Warning: Grader Assignment Count mismatch Submission Count (%d != %d)\n", graderAssignmentCount, submissionCount)
		fmt.Printf("Some graders will (randomly) get more assignments than they requested. To prevent random assignment, increase the 'grade' field for graders in graders.json\n")
	} else if submissionCount < graderAssignmentCount {
		fmt.Printf("Warning: Grader Assignment Count mismatch Submission Count (%d != %d)\n", graderAssignmentCount, submissionCount)
		fmt.Printf("Some graders will (randomly) get fewer assignments than they requested.\n")
	}

	graderList, err := functions.MakeGraderList(graders.G)

	fmt.Printf("\n\n\n")
	printGraderList(graderList)
	fmt.Printf("\n\n\n")
}

func printGraderList(input map[string]*[]string) {
	for grader, gradees := range input {
		for _, g := range *gradees {
			fmt.Printf("%s|%s\n", grader, g)
		}
	}
}