package main

import (
	"fmt"
	"io"

	"github.com/peterh/liner"
	"robpike.io/ivy/mobile"
)

func eval(cmd string) string {
	result, err := mobile.Eval(cmd)
	if err != nil {
		return "Error" + err.Error()
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
			fmt.Println(eval(cmd))
		default:
			fmt.Println(err)
		}
	}
}
