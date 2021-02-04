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
		cmd := exec.Command("./scripts/" + script.Name(), SubmissionsDirectory)
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
		fmt.Printf("%s\n", output)
		fmt.Printf("  %s exited with code %d\n", script.Name(), cmd.ProcessState.ExitCode())
	}

	return nil
}