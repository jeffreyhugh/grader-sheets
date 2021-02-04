package functions

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func RunScripts(SubmissionsDirectory string) error {
	scripts, err := ioutil.ReadDir("./scripts")
	if err != nil {
		return err
	}

	for _, script := range scripts {
		cmd := exec.Command(script.Name(), SubmissionsDirectory)
		fmt.Printf("  %s exited with code %d\n", script.Name(), cmd.ProcessState.ExitCode())
	}

	return nil
}