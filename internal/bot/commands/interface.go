package commands

type CommandManager interface {
	Get() Commands
	GetByName(name string) Command
	GetByRoute(route string) Command
}
