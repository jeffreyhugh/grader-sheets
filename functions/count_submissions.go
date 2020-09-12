package functions

import (
	"io/ioutil"
	"strings"
)

const submissionDirectory = "./submissions"

func CountSubmissions() (int, error) {
	files, err := ioutil.ReadDir(submissionDirectory) // TODO changeable directory
	if err != nil {
		return 0, err
	}

	submissionCount := 0

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".tar.gz") {
			submissionCount++
		}
	}

	return submissionCount, nil
}