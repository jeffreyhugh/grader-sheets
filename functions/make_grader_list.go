package functions

import (
	"../structures"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func MakeGraderList(config *structures.Config, graders *[]structures.Grader) (map[string]*[]string, error) {
	gradersDereferenced := *graders
	graderList := make(map[string]*[]string, 0)
	for _, g := range *graders {
		gradees := make([]string, 0)
		graderList[fmt.Sprintf("%s%s", g.Last, g.First)] = &gradees
	}

	files, err := ioutil.ReadDir(config.SubmissionsDirectory)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(files), func(i, j int) { files[i], files[j] = files[j], files[i] })
	rand.Shuffle(len(gradersDereferenced), func(i, j int) { gradersDereferenced[i], gradersDereferenced[j] = gradersDereferenced[j], gradersDereferenced[i] })

	for _, f := range files {
		author := strings.Split(f.Name(), "_")[0]
		foundGrader := false

		var g structures.Grader
		gradersDereferencedCopy := make([]structures.Grader, len(gradersDereferenced))
		copy(gradersDereferencedCopy, gradersDereferenced)

		// attempt to assign a grader
		if !foundGrader {
			for {
				gradersDereferencedCopy, g = pop(gradersDereferencedCopy)
				if g.Last == "" { // basically if pop failed
					break
				} else if !isConflict(g, author) && len(*graderList[fmt.Sprintf("%s%s", g.Last, g.First)]) < g.Grade { // haha n^3 go brrrt
					*graderList[fmt.Sprintf("%s%s", g.Last, g.First)] = append(*graderList[fmt.Sprintf("%s%s", g.Last, g.First)], f.Name())
					foundGrader = true
					break
				}
			}
		}

		gradersDereferencedCopy = make([]structures.Grader, len(gradersDereferenced))
		copy(gradersDereferencedCopy, gradersDereferenced)
		// assign to random grader without conflict, regardless of number of assignments
		if !foundGrader {
			for {
				gradersDereferencedCopy, g = pop(gradersDereferencedCopy)
				if g.Last == "" {
					break
				} else if !isConflict(g, author) {
					*graderList[fmt.Sprintf("%s%s", g.Last, g.First)] = append(*graderList[fmt.Sprintf("%s%s", g.Last, g.First)], f.Name())
					fmt.Printf("[WARN] assigned %s to %s while ignoring requested workload\n", f.Name(), fmt.Sprintf("%s%s", g.Last, g.First))
					foundGrader = true
					break
				}
			}
		}

		// all graders have a conflict with this assignment
		if !foundGrader {
			fmt.Printf("[WARN] Assignment %s has no valid graders\n", f.Name())
		}
	}

	return graderList, nil
}

func isConflict(g structures.Grader, author string) bool {
	for _, c := range g.Conflicts {
		if author == c {
			return true
		}
	}
	return false
}

func pop(input []structures.Grader) ([]structures.Grader, structures.Grader) {
	inputCopy := input
	if len(input) == 0 {
		return nil, structures.Grader{}
	}
	choice := rand.Intn(len(inputCopy))
	popped := inputCopy[choice]
	inputCopy[choice] = inputCopy[len(inputCopy) - 1]
	return inputCopy[:len(inputCopy)-1], popped
}