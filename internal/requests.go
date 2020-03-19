package internal

// ==================================================
// All request must have this
// ==================================================
type RequestHeaders struct {
	reqid string
	cmd   string
}

// ==================================================
// Request to execute a bash shell command
// ==================================================
type BashRequest struct {
	RequestHeaders
	data []string
}

// ==================================================
// Request to
// + send files to server
// + download files from URLs to HOME
// ==================================================
type FilesRequest struct {
	RequestHeaders
	action string
	files  []string
}
