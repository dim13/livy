package main

import (
	"fmt"
	"io"
	"log"

	"gopkg.in/readline.v1"
	"robpike.io/ivy/mobile"
)

func main() {
	rl, err := readline.New("      ")
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		switch err {
		case io.EOF, readline.ErrInterrupt:
			return
		case nil:
			result, err := mobile.Eval(line)
			if err != nil {
				fmt.Println("Error", err)
				continue
			}
			fmt.Println(result)
		default:
			log.Println(err)
		}
	}
}
