package message

type CommandInfo interface {
	Command() string
	Data() string
}
