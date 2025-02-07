package lexer

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
	for _, ch := range input {
		switch ch {
		case '(':
			tokens = append(tokens, Token{Type: TokenParenL, Value: "("})
		case ')':
			tokens = append(tokens, Token{Type: TokenParenR, Value: ")"})
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			tokens = append(tokens, Token{Type: TokenNumber, Value: string(ch)})
		default:
			tokens = append(tokens, Token{Type: TokenSymbol, Value: string(ch)})
		}
	}
	return tokens
}
