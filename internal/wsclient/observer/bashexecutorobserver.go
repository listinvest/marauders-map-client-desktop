package observer

import (
	"log"
	"os/exec"
)

// Bash executor commands Observer
type BashExecutorObserver struct {
}

// The first param is the command interpreted by the client
// not the bash command or os program
// The data variable, is an array of all bash command (or program) data
// that must be Joined into one string command
// Eg:
// "ls -l -a"
// "rm -rfv directory/"
func (o *BashExecutorObserver) execute(cmd string, data []string) {
	if cmd != "bash" {
		return
	}

	if len(data) <= 0 {
		return
	}

	log.Println("BashExecutorObserver: new action triggered")

	// Invokes command (or program)
	o.executeCommand(data)
}

// Executes a command (or program) directly with it's params
// @scmd:
func (o *BashExecutorObserver) executeCommand(scmd []string) string {
	cmd := scmd[0]        // First position es the command (or program). Eg: ls, rm, mkdir
	cmdparams := scmd[1:] // Command (or program) params

	cmdres := exec.Command(cmd, cmdparams...)
	stdout, err := cmdres.Output()

	if err != nil {
		return ""
	}

	res := string(stdout)
	return res
}

func NewBashExecutorObserver() *BashExecutorObserver {
	return &BashExecutorObserver{}
}
