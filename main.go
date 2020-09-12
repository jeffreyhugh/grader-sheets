package main

import (
	"./functions"
	"./structures"
	"fmt"
	"strings"
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
	printGraderEmails(graders)
	fmt.Printf("\n")
	printGraderList(graderList)
	fmt.Printf("\n\n\n")
}

func printGraderList(input map[string]*[]string) {
	fmt.Printf("Name (LastFirst)|Student|Filename\n")
	for grader, gradees := range input {
		for _, g := range *gradees {
			fmt.Printf("%s|%s|%s\n", grader, strings.Split(g, "_")[0], g)
		}
	}
}

func printGraderEmails(input *structures.Graders) {
	fmt.Printf("Name (LastFirst)|Email\n")
	for _, g := range *input.G {
		fmt.Printf("%s%s|%s\n", g.Last, g.First, g.Email)
	}
}