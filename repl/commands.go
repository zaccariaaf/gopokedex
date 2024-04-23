package repl

type command struct {
	name     string
	usage    string
	callback func() error
}
