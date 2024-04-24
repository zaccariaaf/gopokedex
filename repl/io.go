package repl

import (
	"bufio"
	"fmt"
	"os"
)

const (
	BASE_PROMPT = "pokedex"
)

func RunRepl(prompt string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	printPrompt(prompt)
	return readToTokens(scanner)
}

func printPrompt(prompt string) {
	fmt.Print("%v> ", prompt)
}

func printMessage(message string) {
	fmt.Printf(message)
}

func printCommandHelpText(command, usage string) {
	fmt.Println("%v:	%v", command, usage)
}

func readToTokens(scanner *bufio.Scanner) []string {
	tokens := []string{}
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return tokens
}
