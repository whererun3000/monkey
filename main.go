package main

import (
	"fmt"
	"os"

	"github.com/whererun3000/monkey/repl"
)

func main() {
	fmt.Printf("Welcome to Monkey\n")
	repl.Start(os.Stdin, os.Stdout)
}
