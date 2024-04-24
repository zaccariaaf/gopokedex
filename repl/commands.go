package repl

import (
	"errors"
	"os"
)

type command struct {
	name     string
	usage    string
	Callback func() error
}

func Parse(tokens []string) ([]command, error) {
	validCommands := getValidCommands()
	var commands []command
	for _, token := range tokens {
		command, ok := validCommands[token]
		if !ok {
			printErrorMessage("invalid command:", token)
			printHelpMesage(validCommands)
			return nil, errors.New("invalid command")
		}
		commands = append(commands, command)
	}
	return commands, nil
}

func getValidCommands() map[string]command {
	return map[string]command{
		"help": {
			name:     "help",
			usage:    "Displays this help message",
			Callback: commandHelp,
		},
		"exit": {
			name:     "exit",
			usage:    "Exit the pokedex",
			Callback: commandExit,
		},
	}
}

func commandHelp() error {
	printMessage("\nWelcome to the Pokedex!\nUsage:")
	printHelpMesage(getValidCommands())
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
