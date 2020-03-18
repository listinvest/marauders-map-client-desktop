package command

// ===========================================
// Command Interface
// Define Command implementations for answering to the server
// ===========================================
type ResponseCommand interface {
	execute()
}
