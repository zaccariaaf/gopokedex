package repl

import (
	"bufio"
)

func readToTokens(scanner *bufio.Scanner) []string {
	tokens := []string{}
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}
	return tokens
}
