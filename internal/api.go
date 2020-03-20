package internal

// ==================================================
// All request must have this
// ==================================================
type RequestHeaders struct {
	Reqid string
	Cmd   string
}

// ==================================================
// Request to execute a bash shell command
// ==================================================
type BashRequest struct {
	RequestHeaders
	Data []string
}

// ==================================================
// Request to
// + send files to server
// + download files from URLs to HOME
// ==================================================
type FilesRequest struct {
	RequestHeaders
	Action string
	Files  []string
}
