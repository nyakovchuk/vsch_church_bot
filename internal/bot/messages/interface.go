package messages

type CommandInfo interface {
	Command() string
	Data() string
}
