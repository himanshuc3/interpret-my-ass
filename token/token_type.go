package token

// NOTE:
// 1. token type is declared as string instead of int/byte which
// might be better for performance and provide space gains (changed it to int)
// 2. Defining aliases can act as enums, provide semantic meaning and reduce the
// scope of accepted characters (because prevents mixing)
// 3. This isn't type alias but a new type
type TokenType int

// NOTE:
// 1. const ILLEGAL string = "ILLEGAL" - explicit const
// Btw, very og syntax with let and const. Well, ackshually
// technically, it was included in ES6 spec in JS, so it might
// have been copied from golang to js
const (
	// Single-character tokens
	TokenType_LEFT_PAREN TokenType = iota
	TokenType_RIGHT_PAREN
	TokenType_LEFT_BRACE
	TokenType_RIGHT_BRACE
	TokenType_COMMA
	TokenType_DOT
	TokenType_MINUS
	TokenType_PLUS
	TokenType_SEMICOLON
	TokenType_SLASH
	TokenType_STAR

	// One or two character tokens
	TokenType_BANG
	TokenType_BANG_EQUAL
	TokenType_EQUAL
	TokenType_EQUAL_EQUAL
	TokenType_GREATER
	TokenType_GREATER_EQUAL
	TokenType_LESS
	TokenType_LESS_EQUAL

	// Literals
	TokenType_IDENTIFIER
	TokenType_STRING
	TokenType_NUMBER

	// Keywords
	TokenType_AND
	TokenType_CLASS
	TokenType_ELSE
	TokenType_FALSE
	TokenType_FOR
	TokenType_FUN
	TokenType_IF
	TokenType_NIL
	TokenType_OR
	TokenType_PRINT
	TokenType_RETURN
	TokenType_SUPER
	TokenType_THIS
	TokenType_TRUE

	TokenType_EOF
)

// NOTE:
// 1. In go, a const can only be a compile time constant like
// number string, boolean
var TokenMapping map[TokenType]Token = GenerateTokenMapping()

// NOTE:
// 1. Can be removed in favor of a global variable since global
// map constants aren't allowed in go
func GenerateTokenMapping() map[TokenType]Token {

	// NOTE:
	// 1. Dynamic property names like in JS [COMMA]
	tokens := map[TokenType]Token{
		ASSIGN: {
			Type: ASSIGN,
		},
		SEMICOLON: {
			Type: SEMICOLON,
		},
		LPAREN: {
			Type: LPAREN,
		},
		COMMA: {
			Type: COMMA,
		},
		PLUS: {
			Type: PLUS,
		},
		LBRACE: {
			Type: LBRACE,
		},
		RBRACE: {
			Type: RBRACE,
		},
	}
	return tokens
}
