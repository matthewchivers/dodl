package core

type Command interface {
	Execute() error
}
