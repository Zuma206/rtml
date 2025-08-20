package streval

import "strings"

type Handler func(str string) error

type Parser struct {
	Token      string
	EscapeChar *byte
	Handler    Handler
	Next       *Parser
}

func (parser *Parser) isEscaped(str string, startOfToken int) (int, bool) {
	potentialEscapeIndex := startOfToken - 1
	potentialEscape := str[potentialEscapeIndex]
	isTokenAtStart := startOfToken == 0
	isTokenEscaped := parser.EscapeChar != nil && !isTokenAtStart && potentialEscape == *parser.EscapeChar
	return potentialEscapeIndex, isTokenEscaped
}

func (parser *Parser) Parse(str string) error {
	startOfToken := strings.Index(str, parser.Token)
	if startOfToken == -1 {
		return parser.Handler(str)
	}
	afterToken := str[startOfToken+len(parser.Token):]
	if startOfEscape, isEscaped := parser.isEscaped(str, startOfToken); isEscaped {
		beforeEscape := str[:startOfEscape]
		if err := parser.Handler(beforeEscape + parser.Token); err != nil {
			return err
		}
		return parser.Parse(afterToken)
	}
	beforeToken := str[:startOfToken]
	if err := parser.Handler(beforeToken); err != nil {
		return err
	}
	return parser.Next.Parse(afterToken)
}

type Handlers struct {
	Literal    Handler
	Expression Handler
}

var escapeChar = byte('\\')

func Parse(str string, handlers Handlers) error {
	literalParser := &Parser{
		Token:      "${",
		EscapeChar: &escapeChar,
		Handler:    handlers.Literal,
	}
	expressionParser := &Parser{
		Token:   "}",
		Handler: handlers.Expression,
		Next:    literalParser,
	}
	literalParser.Next = expressionParser
	return literalParser.Parse(str)
}
