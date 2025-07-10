package main

import (
	"fmt"
	"os"

	"github.com/himanshuc3/interpret-my-ass/lang"
)

// NOTE:
// 1. Support REPL and file execution
func main() {
	// NOTE:
	// 1. Automatic conversion from struct to pointer in interface methods
	// Yes, Golang is smart and shitty concurrently
	l := lang.GetInterpreter()

	if len(os.Args) > 2 {
		fmt.Println("Usage: goalang [script]")
		// NOTE:
		// 1. 64 - CLI usage error
		os.Exit(64)
	}

	if len(os.Args) == 2 {
		l.RunFile(os.Args[1])
	} else {
		l.RunPrompt()
	}
}
