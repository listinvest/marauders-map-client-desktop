package internal

import (
	"os/exec"
)

// Bash executor commands Observer
type BashExecutorObserver struct {
}

func (o *BashExecutorObserver) execute(string_json string) {

	// bashres := o.executeCommand(req.Data)
	// bashres.Reqid = req.Reqid

}

// Executes a command (or program) directly with it's params
// @scmd: array of program and it's params for execution
func (o *BashExecutorObserver) executeCommand(scmd []string) BashResponse {
	program := scmd[0]        // First position es the command (or program). Eg: ls, rm, mkdir
	programparams := scmd[1:] // Command (or program) params

	cmdres := exec.Command(program, programparams...)
	stdout, err := cmdres.Output()

	// Prepare response
	bashres := BashResponse{}

	if err != nil {
		bashres.Err = true
		bashres.Errmsg = err.Error()
		return bashres
	}

	bashres.Err = false
	bashres.Result = string(stdout)

	// TODO: return 'BashResponse' variable
	return bashres
}

func NewBashExecutorObserver() *BashExecutorObserver {
	return &BashExecutorObserver{}
}
