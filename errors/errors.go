package errors

import (
	err "errors"
	"fmt"
)

var (
	New = err.New
)

var ErrUnterminatedString = New("Unterminated string")
var ErrInvalidNumberLiteral = New("invalid number literal")

func ErrUnexpectedCharacter(c rune) error {
	return fmt.Errorf("unexpected character %c", c)
}
