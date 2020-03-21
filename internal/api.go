package internal

// ==================================================
// All requests must have this
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

// ==================================================
// Request to
// + start record screen
// + stop record screen
// + take screenshot
// ==================================================
type ScreenRequest struct {
	RequestHeaders
	Action        string
	Action_status string
	Seconds       int
}

// ==========================================================================================

// ==================================================
// All responses must have this
// ==================================================
type ResponseHeaders struct {
	Reqid  string `json:"reqid"`
	Err    bool   `json:"err"`
	Errmsg string `json:"errmsg"`
}

// ==================================================
// Response of shell command execution
// ==================================================
type BashResponse struct {
	ResponseHeaders
	Result string `json:"result"`
}

// ==================================================
// Response of shell command execution
// ==================================================
type ScreenshotNotification struct {
	ResponseHeaders
	Filename string `json:"filename"`
}
