package main

import (
	"fmt"
	"os"
	"os/user"
	"strange/repl"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Strange programming language!\n", currentUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
