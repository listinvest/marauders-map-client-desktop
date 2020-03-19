package internal

// ===========================================
// Command Interface
// Define Command implementations for answering to the server
// ===========================================
type ResponseCommand interface {
	execute()
}

type SendFileCommand struct {
	AbsolutePath string
	wsc          *WSClient
}
