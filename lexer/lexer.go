package lexer

import (
	"fmt"

	"github.com/himanshuc3/interpret-my-ass/token"
)

// NOTE: Type interconversions are important: byte - string - TokenType (string)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NOTE:
// 1. Supporting only ASCII instead of UTF8 to prevent
// complexity, also keeps ch limited to byte
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII for NUL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// TODO: Standardize nul parsing in tokenmapping

	if l.ch == 0 {
		tok.Literal = ""
		tok.Type = token.EOF
		l.readChar()
		return tok
	}

	newTokenMapping, ok := token.TokenMapping[token.TokenType(l.ch)]

	if !ok {
		fmt.Printf("Invalid token processed: %v", token.TokenType(l.ch))
		panic("Invalid token processed")
	}

	tok = newTokenMapping
	tok.Literal = string(l.ch)

	l.readChar()

	return tok
}

// TODO: Not required currently because of the token mapping specified
// func newToken(tokenType token.TokenType, ch byte) token.Token {

// }
