package main

import (
	"fmt"

	"github.com/peterh/liner"
	"robpike.io/ivy/mobile"
)

const prompt = "      "

func main() {
	l := liner.NewLiner()
	defer l.Close()
	for {
		line, err := l.Prompt(prompt)
		if err != nil {
			fmt.Println(err)
			break
		}
		if line == ")off" {
			break
		}
		l.AppendHistory(line)
		result, err := mobile.Eval(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
