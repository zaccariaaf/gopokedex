package repl

type command struct {
	name     string
	usage    string
	callback func() error
}

func GetValidCommands() map[string]command {
	return map[string]command{
		"help": {
			name:     "help",
			usage:    "Displays this help message",
			callback: commandHelp,
		},
		"exit": {
			name:     "exit",
			usage:    "Exit the pokedex",
			callback: commandExit,
		},
	}
}

func commandHelp() error {
	printMessage("\nWelcome to the Pokedex!\nUsage:")
	for _, command := range GetValidCommands() {
		printCommandHelpText(command.name, command.usage)
	}
	return nil
}

func commandExit() error {
	return nil
}
