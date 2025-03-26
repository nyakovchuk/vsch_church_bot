package logmessage

type CommandInfo interface {
	Command() string
	Data() string
}
