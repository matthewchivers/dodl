package core

// Command is an interface for any command that can be executed.
type Command interface {
	Execute() error
}
