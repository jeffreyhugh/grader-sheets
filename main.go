package main

import (
	"./functions"
	"./structures"
	"bufio"
	"fmt"
	"os"
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
		fmt.Printf("Warning: Grader Assignment Count less than Submission Count (%d < %d)\n", graderAssignmentCount, submissionCount)
		fmt.Printf("Some graders will (randomly) get more assignments than they requested. To prevent random assignment, increase the 'grade' field for graders in graders.json\n")
	} else if submissionCount < graderAssignmentCount {
		fmt.Printf("Warning: Grader Assignment Count greater than Submission Count (%d > %d)\n", graderAssignmentCount, submissionCount)
		fmt.Printf("Some graders will (randomly) get fewer assignments than they requested.\n")
	}

	graderList, err := functions.MakeGraderList(graders.G)

	printGraderEmails(graders)
	printGraderList(graderList)
	fmt.Printf("\nSuccessfully wrote grader info and grader sheet. Run ./print.sh to view the results.\n")
}

func printGraderList(input map[string]*[]string) {
	file, err := os.Create("./graderlist.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	_, err = fmt.Fprintf(w, "Name (LastFirst)|Student|Filename\n")
	if err != nil {
		panic(err)
	}
	for grader, gradees := range input {
		for _, g := range *gradees {
			_, err := fmt.Fprintf(w, "%s|%s|%s\n", grader, strings.Split(g, "_")[0], g)
			if err != nil {
				panic(err)
			}
		}
	}

	w.Flush()
}

func printGraderEmails(input *structures.Graders) {
	file, err := os.Create("./graderemails.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	_, err = fmt.Fprintf(w, "Name (LastFirst)|Email\n")
	if err != nil {
		panic(err)
	}
	for _, g := range *input.G {
		_, err := fmt.Fprintf(w, "%s%s|%s\n", g.Last, g.First, g.Email)
		if err != nil {
			panic(err)
		}
	}

	w.Flush()
}
