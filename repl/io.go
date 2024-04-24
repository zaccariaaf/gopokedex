package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	BasePrompt = "pokedex"
)

func GetInput(prompt string, tokenChannel chan []string) {
	scanner := bufio.NewScanner(os.Stdin)
	printPrompt(prompt)
	nextLine := readNextLine(scanner)
	tokenChannel <- strings.Split(nextLine, " ")
}

func printPrompt(prompt string) {
	fmt.Printf("%v> ", prompt)
}

func printMessage(message string) {
	fmt.Print(message)
}

func printErrorMessage(message string, args ...string) {
	fmt.Fprintln(os.Stderr, message, args)
}

func printHelpMesage(validCommands map[string]command) {
	for _, command := range validCommands {
		printCommandHelpText(command.name, command.usage)
	}
}

func printCommandHelpText(command, usage string) {
	fmt.Printf("%v:	%v\n", command, usage)
}

func readNextLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return scanner.Text()
}
