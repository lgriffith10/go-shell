package parser

import (
	"strings"
)

type EscapeMode int

const (
	EscapeNone EscapeMode = iota
	EscapeSingleQuote
	EscapeDoubleQuote
	EscapeBackslash
)

func ParseCommand(input string) []string {
	var args []string
	var current strings.Builder

	escaped := false
	quoteMode := EscapeNone

	for i := 0; i < len(input); i++ {
		c := input[i]

		if escaped {
			if quoteMode == EscapeDoubleQuote &&
				c != '\\' && c != '"' && c != '$' && c != '`' {
				current.WriteByte('\\')
			}
			current.WriteByte(c)
			escaped = false
			continue
		}

		if c == '\\' {
			escaped = true
			continue
		}

		if quoteMode != EscapeSingleQuote {
			if c == '"' {
				switch quoteMode {
				case EscapeDoubleQuote:
					quoteMode = EscapeNone
				case EscapeNone:
					quoteMode = EscapeDoubleQuote
				}
				continue
			}
		}

		if quoteMode != EscapeDoubleQuote {
			if c == '\'' {
				switch quoteMode {
				case EscapeSingleQuote:
					quoteMode = EscapeNone
				case EscapeNone:
					quoteMode = EscapeSingleQuote
				}
				continue
			}
		}

		if c == ' ' && quoteMode == EscapeNone {
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
			continue
		}

		current.WriteByte(c)
	}

	if current.Len() > 0 {
		args = addArg(args, current)
	}

	return args
}

func addArg(args []string, current strings.Builder) []string {
	args = append(args, current.String())
	current.Reset()
	return args
}
