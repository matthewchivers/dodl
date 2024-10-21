package models

type CommandContext struct {
	Command    string
	Args       []string
	Flags      map[string]interface{}
	EntryPoint string
}
