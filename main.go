package main

import (
	"bufio"
	"fmt"
	"os"
)

type GoaLang struct {
	hadError bool // Task: Really bad naming, tells the status
}

// NOTE:
// 1. Support REPL and file execution

func main() {
	// NOTE:
	// 1. Automatic conversion from struct to pointer in interface methods
	// Yes, Golang is smart and shitty concurrently
	l := GoaLang{}

	if len(os.Args) > 2 {
		fmt.Println("Usage: goalang [script]")
		// NOTE:
		// 1. 64 - CLI usage error
		os.Exit(64)
	}

	if len(os.Args) == 2 {
		l.runFile(os.Args[1])
	} else {
		l.runPrompt()
	}
}

func (l *GoaLang) runPrompt() error {
	// NOTE:
	// 1. Reading from stdin
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if ok := input.Scan(); !ok {
			break
		}
		line := input.Text()
		l.run(line)
		l.hadError = false
	}
}

func (l *GoaLang) runFile(path string) error {
	bytes, err := os.ReadFile(path)

	if err != nil {
		return nil
	}
	l.run(string(bytes))
	if l.hadError {
		os.Exit(65)
	}
	return nil
}

func (l *GoaLang) run(source string) error {
	// Task: Get the source code, parse it into tokens and print them
	scanner := Scanner{source}
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func (l *GoaLang) reportError(line int, msg string) {
	l.report(line, "", msg)
}

func (l *GoaLang) report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
	l.hadError = true
}
