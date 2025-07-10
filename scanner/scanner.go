package scanner

import (
	"errors"
	"fmt"

	"github.com/himanshuc3/interpret-my-ass/token"
)

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
		// Task: skip a comment line
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
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
		s.hadError = true
		s.report(s.line, "", fmt.Sprintf("Unexpected character: %c", c))

	}
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
		s.report(s.line, "", "Unterminated string.")
		return
	}

	// The closing "
	s.advance()

	value := string(s.source[s.start+1 : s.current-1])
	s.addTokenWithLiteral(token.TokenType_STRING, value)
}

func (s *Scanner) ReportError(line int, msg string) {
	s.report(line, "", msg)
}

// NOTE:
// 1. Error reporting should be the first thing to handle
// given it acts as the first line of defense against hours of debugging
func (s *Scanner) report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
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
