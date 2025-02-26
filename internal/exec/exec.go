package main

import (
	"fmt"
	"go_lisp_interpreter/internal/eval"
	"go_lisp_interpreter/internal/lexer"
	"go_lisp_interpreter/internal/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lisp <script.lisp>")
		return
	}
	filename := os.Args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	env := eval.NewEnv(nil)
	tokens := lexer.Tokenize(string(content))
	ast, _ := parser.Parse(tokens)
	result := eval.Eval(ast, env)
	fmt.Println(result)
}
