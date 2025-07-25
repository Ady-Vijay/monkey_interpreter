package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, this is the monkey programming language\n", user.Username)
	fmt.Printf("Feel free to type commands")

	repl.Start(os.Stdin, os.Stdout)
}
