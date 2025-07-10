package lang

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/himanshuc3/interpret-my-ass/scanner"
)

// NOTE:
// 1. Circular deps are a major pain-point when dealing with BE modelling
type GoaLang struct {
	// hadError bool // Task: Really bad naming, tells the status (shifted to error package)
}

var (
	// NOTE:
	// 1. Always prefer minimal exposure, like a good old
	// conservative indian family
	// So, even though the variable is private to the package,
	// we can define a public getter
	interpreter *GoaLang
	once        sync.Once
)

func GetInterpreter() *GoaLang {
	// NOTE:
	// 1. There is a difference between sync.Once and init (instance vs package level, thread safety etc.)
	once.Do(func() {
		fmt.Println("Creating interpreter fr fr")
		interpreter = &GoaLang{
			// hadError: false,
		}
	})
	return interpreter
}

func (l *GoaLang) RunPrompt() error {
	// NOTE:
	// 1. Reading from stdin
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if ok := input.Scan(); !ok {
			break
		}
		line := input.Text()
		l.Run(line)
		// l.hadError = false
	}
	return nil
}

func (l *GoaLang) RunFile(path string) error {
	bytes, err := os.ReadFile(path)

	if err != nil {
		return nil
	}
	err = l.Run(string(bytes))
	if err != nil {
		os.Exit(65)
	}
	return nil
}

func (l *GoaLang) Run(source string) error {
	// Task: Get the source code, parse it into tokens and print them
	scanObj := scanner.NewScanner(source)
	err := scanObj.ScanTokens()

	if err != nil {
		return err
	}

	// NOTE:
	// 1. Golang - no, Golund - yes. Even the getters returning the array of
	// pointers can be edited
	tokens := scanObj.GetTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
	return nil
}
