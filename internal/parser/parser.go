package parser

import "go_lisp_interpreter/internal/lexer"

type Node struct {
	Type     string
	Value    string
	Children []Node
}

func Parse(tokens []lexer.Token) (Node, int) {
	if len(tokens) == 0 {
		panic("Empty token list")
	}
	return parseExpr(tokens, 0)
}

func parseExpr(tokens []lexer.Token, pos int) (Node, int) {
	if tokens[pos].Type == lexer.TokenParenL {
		children := []Node{}
		pos++
		for pos < len(tokens) && tokens[pos].Type != lexer.TokenParenR {
			child, newPos := parseExpr(tokens, pos)
			children = append(children, child)
			pos = newPos
		}
		return Node{Type: "LIST", Children: children}, pos + 1
	}
	return Node{Type: "ATOM", Value: tokens[pos].Value}, pos + 1
}
