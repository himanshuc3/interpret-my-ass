package token

import "fmt"

// NOTE:
// 1. What is the difference between lexel and object?
// 2. Go stringer tool - helps autogenerate Stringer interface on structs
type Token struct {
	TokenType TokenType
	Lexeme    string
	Object    any
	Line      int
}

// NOTE:
// 1. Factory pattern: Commonly used as constructor for
// objects that require initialization
func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{
		tokenType, lexeme, literal, line,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %v", t.TokenType, t.Lexeme, t.Object)
}
