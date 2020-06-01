package commands

// The Response value from Execute.
type Response struct {
	// The build Result will only be set in the hugo build command.
	Result string

	// Err is set when the command failed to execute.
	Err error

	// The command that was executed.
	Cmd string
}

// Execute against the REST API
func Execute(args []string) Response {
	var resp Response

	return resp
}
