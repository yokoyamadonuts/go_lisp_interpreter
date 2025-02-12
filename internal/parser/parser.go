package parser

import "go_lisp_interpreter/internal/lexer"

type Node struct {
	Type     string
	Value    string
	Children []Node
}

func Parse(tokens []lexer.Token) Node {
	return Node{Type: "root", Children: []Node{}}
}
