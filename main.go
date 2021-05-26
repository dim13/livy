package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
	"robpike.io/ivy/mobile"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	t := term.NewTerminal(os.Stdin, "      ")
	for {
		line, err := t.ReadLine()
		if err != nil {
			fmt.Fprintln(t, err)
			break
		}
		if line == ")off" {
			break
		}
		result, err := mobile.Eval(line)
		if err != nil {
			fmt.Fprintln(t, err)
			continue
		}
		fmt.Fprintln(t, result)
	}
}
