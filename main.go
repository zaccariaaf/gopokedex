package main

import (
	"github.com/zaccariaaf/gopokedex/repl"
)

func main() {
	for {
		ch := make(chan []string)
		go repl.GetInput(repl.BasePrompt, ch)
		tokens := <-ch
		commands, err := repl.Parse(tokens)
		if err != nil {
			break
		}
		for _, command := range commands {
			err := command.Callback()
			if err != nil {
				break
			}
		}
	}
}
