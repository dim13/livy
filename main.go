package main

import (
	"fmt"
	"io"

	"github.com/peterh/liner"
	"robpike.io/ivy/mobile"
)

func eval(cmd string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Sprintf("Panic %v", r)
		}
	}()
	result, err := mobile.Eval(cmd)
	if err != nil {
		return fmt.Sprintf("Error %v", err)
	}
	return result
}

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)
	for {
		cmd, err := line.Prompt("      ")
		switch err {
		case liner.ErrPromptAborted, io.EOF:
			return
		case nil:
			line.AppendHistory(cmd)
			fmt.Print(eval(cmd))
		default:
			fmt.Print(err)
		}
	}
}
