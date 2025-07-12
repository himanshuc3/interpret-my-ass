package scanner

import (
	"fmt"
	"strconv"

	"github.com/himanshuc3/interpret-my-ass/errors"

	"github.com/himanshuc3/interpret-my-ass/token"
)

// NOTE:
// 1. Nuances handling newlines in tokenization parsing
// 2. Chomsky heirarchy - parsing and AST

type Scanner struct {
	// Task: replace string with []rune to have a more
	// general implementation of the interpreter
	source []rune
	tokens []*token.Token

	start    int
	current  int
	line     int
	hadError bool
}

func (s *Scanner) GetTokens() []*token.Token {
	return s.tokens
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: []rune(source),
		tokens: make([]*token.Token, 0),
		line:   1,
	}
}

func (s *Scanner) ScanTokens() error {
	for !s.isAtEnd() {
		// Task: Move s.start assignment to scanning tokens
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, &token.Token{
		TokenType: token.TokenType_EOF,
		Lexeme:    "",
		Object:    nil,
		Line:      s.line,
	})

	if s.hadError {
		return errors.New("Error in tokenization") // Needs to be collated list of messages later
	} else {
		return nil
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	var c rune = s.advance()

	// Task:
	// Can probably replace this with a map or regex match
	switch c {
	// Task: Replace strings with const tokens
	// Possibly create mappings as well, but golang doesn't
	// like abstracting code
	case '(':
		s.addToken(token.TokenType_LEFT_PAREN)
	case ')':
		s.addToken(token.TokenType_RIGHT_PAREN)
	case '{':
		s.addToken(token.TokenType_LEFT_BRACE)
	case '}':
		s.addToken(token.TokenType_RIGHT_BRACE)
	case ',':
		s.addToken(token.TokenType_COMMA)
	case '.':
		s.addToken(token.TokenType_DOT)
	case '-':
		s.addToken(token.TokenType_MINUS)
	case '+':
		s.addToken(token.TokenType_PLUS)
	case ';':
		s.addToken(token.TokenType_SEMICOLON)
	case '*':
		s.addToken(token.TokenType_STAR)
	case '!':
		if s.match('=') {
			s.addToken(token.TokenType_BANG_EQUAL)
		} else {
			s.addToken(token.TokenType_BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.TokenType_EQUAL_EQUAL)
		} else {
			s.addToken(token.TokenType_EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.TokenType_LESS_EQUAL)
		} else {
			s.addToken(token.TokenType_LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.TokenType_GREATER_EQUAL)
		} else {
			s.addToken(token.TokenType_GREATER)
		}
	case '/':

		if s.match('/') {
			s.consumeLineComment()
		} else if s.match('*') {
			s.consumeMultiLineComment()
		} else {
			s.addToken(token.TokenType_SLASH)
		}
	case ' ', '\r', '\t':
		// Task: handle characters

	case '\n':
		s.line++
	case '"':
		s.scanString()

	default:
		if isDigit(c) {
			s.scanNumber()
		} else if isAlpha(c) {
			s.scanIdentifier()
		} else {
			s.hadError = true
			s.report(s.line, "", errors.ErrUnexpectedCharacter(c))
		}

	}
}

func (s *Scanner) consumeLineComment() {
	for s.peek() != '\n' && !s.isAtEnd() {
		s.advance()
	}
}

func (s *Scanner) consumeMultiLineComment() {

	for !s.isAtEnd() {

		switch c := s.peek(); c {
		case '\n':
			s.line++
			s.advance()
		case '*':
			if s.peekNext() == '/' {
				s.current += 2
				return
			}
		default:
			s.advance()
		}

	}
	s.ReportError(s.line, errors.ErrUnterminatedMultilineComment)

}

func (s *Scanner) scanIdentifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := string(s.source[s.start:s.current])
	if t, ok := token.IsKeyword(text); ok {
		s.addToken(t)
		return
	}
	s.addToken(token.TokenType_IDENTIFIER)
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c rune) bool {
	return isAlpha(c) || isDigit(c)
}

func (s *Scanner) scanString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.hadError = true
		s.report(s.line, "", errors.ErrUnterminatedString)
		return
	}

	// The closing "
	s.advance()

	value := string(s.source[s.start+1 : s.current-1])
	s.addTokenWithLiteral(token.TokenType_STRING, value)
}

func (s *Scanner) ReportError(line int, err error) {
	s.report(line, "", err)
}

// NOTE:
// 1. Error reporting should be the first thing to handle
// given it acts as the first line of defense against hours of debugging
func (s *Scanner) report(line int, where string, err error) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, err)
	s.hadError = true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}

	return s.source[s.current]
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) advance() rune {
	r := s.source[s.current]
	s.current++
	return r
}

func (s *Scanner) addToken(tokenType token.TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType token.TokenType, literal any) {
	s.tokens = append(s.tokens, &token.Token{
		// NOTE:
		// 1. Too explicit, even though tokentype implicitly
		// gives us the lexeme, we store it as well
		TokenType: tokenType,
		Lexeme:    string(s.source[s.start:s.current]),
		Object:    literal,
		Line:      s.line,
	})
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) scanNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	text := string(s.source[s.start:s.current])
	// Task: Report a parsing error if an error with parsing number
	num, err := strconv.ParseFloat(text, 64)
	if err != nil {
		// Task: Use the built-in error interfrace
		// instead of plain string literals
		s.ReportError(s.line, errors.ErrInvalidNumberLiteral)
	}
	s.addTokenWithLiteral(token.TokenType_NUMBER, num)

}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '\000'
	}

	return s.source[s.current+1]
}
