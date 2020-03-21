package internal

import (
	"encoding/json"
	"log"
	"os/exec"
)

// Bash executor commands Observer
type BashExecutorObserver struct {
	respondServerCmd *RespondServerCommand
}

func (o *BashExecutorObserver) execute(string_json string) {
	var req BashRequest
	err := json.Unmarshal([]byte(string_json), &req)

	if err != nil {
		log.Println("ERRROR Unmarshing: ", err)
		return
	}

	if req.Cmd != "bash" {
		return
	}

	log.Println("BashExecutorObserver: received: ", string_json)

	if len(req.Data) <= 0 {
		return
	}

	bashres := o.executeCommand(req.Data)
	bashres.Reqid = req.Reqid

	errr := o.respondServerCmd.SendBashResponse(bashres)

	// TODO: delete this
	if errr == nil {
		strres, _ := json.Marshal(bashres)
		log.Println("BashExecutorObserver: responded: ", string(strres))
	}
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

func NewBashExecutorObserver(respondServerCmd *RespondServerCommand) *BashExecutorObserver {
	return &BashExecutorObserver{
		respondServerCmd: respondServerCmd,
	}
}
