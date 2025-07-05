package lexer

import (
	"testing"

	"github.com/himanshuc3/interpret-my-ass/token"
)

// NOTE:
// 1. All files with *_test.go extension and test names starting with Test*
func TestNextToken(t *testing.T) {

	input := `=+(){},;`

	// NOTE:
	// 1. Expected results from Parents
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		// NOTE:
		// 1. Me giving the wrong results everytime without fail
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			// NOTE:
			// 1. Fatalf - Errorf + Abort test + good for fail stack trace
			// 2. %d - decimal vs %q - string with quotes "foo"
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
