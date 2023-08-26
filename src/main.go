package main

import (
	"fmt"
	"lox/repl"
	"os"
)

func main() {
	fmt.Println("LOX Version 1.0")
	fmt.Println("Enter 'help' to view available commands")

	repl.Start(os.Stdin, os.Stdout)
}
