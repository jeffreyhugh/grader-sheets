package functions

import (
	"../external/unzip"
	"../structures"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func CountSubmissions(config *structures.Config) (int, error) {
	submissionsDirectory, err := unzipSubmissions(config.SubmissionsZip)
	if err != nil {
		return 0, err
	}

	config.SubmissionsDirectory = submissionsDirectory

	files, err := ioutil.ReadDir(submissionsDirectory)
	if err != nil {
		return 0, err
	}

	submissionCount := 0

	for _, f := range files {
		if strings.HasSuffix(f.Name(), config.FileType) {
			submissionCount++
		}
	}

	return submissionCount, nil
}

func unzipSubmissions(submissionsZip string) (string, error) {
	unzipDestination := fmt.Sprintf("submissions-%s", time.Now().Format("2006-01-02-15-04-05"))
	if err := unzip.New(submissionsZip, unzipDestination).Extract(); err != nil {
		return "", err
	}

	return unzipDestination, nil
}
