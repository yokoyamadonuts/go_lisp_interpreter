package lexer

import (
	"strings"
	"unicode"
)

type TokenType string

const (
	TokenNumber TokenType = "NUMBER"
	TokenSymbol TokenType = "SYMBOL"
	TokenParenL TokenType = "PAREN_L"
	TokenParenR TokenType = "PAREN_R"
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(input string) []Token {
	var tokens []Token
	input = strings.TrimSpace(input)
	i := 0

	for i < len(input) {
		switch ch := input[i]; {
		case ch == '(':
			tokens = append(tokens, Token{TokenParenL, "("})
			i++
		case ch == ')':
			tokens = append(tokens, Token{TokenParenR, ")"})
			i++
		case unicode.IsDigit(rune(ch)):
			start := i
			for i < len(input) && unicode.IsDigit(rune(input[i])) {
				i++
			}
			tokens = append(tokens, Token{TokenNumber, input[start:i]})
		case unicode.IsLetter(rune(ch)):
			start := i
			for i < len(input) && (unicode.IsLetter(rune(input[i])) || unicode.IsDigit(rune(input[i]))) {
				i++
			}
			tokens = append(tokens, Token{TokenSymbol, input[start:i]})
		case unicode.IsSpace(rune(ch)):
			i++
		default:
			panic("Unexpected character: " + string(ch))
		}
	}
	return tokens
}
