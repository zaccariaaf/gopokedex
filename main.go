package main

import (
	"fmt"

	"github.com/zaccariaaf/gopokedex/repl"
)

func main() {
	for {
		input := repl.RunRepl(repl.BASE_PROMPT)
		fmt.Print(input)
	}
}
