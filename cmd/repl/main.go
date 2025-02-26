package main

import (
	"bufio"
	"fmt"
	"go_lisp_interpreter/internal/eval"
	"go_lisp_interpreter/internal/lexer"
	"go_lisp_interpreter/internal/parser"
	"os"
)

func main() {
	env := eval.NewEnv(nil)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Lisp REPL. Type 'exit' to quit.")
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			break
		}
		tokens := lexer.Tokenize(text)
		ast, _ := parser.Parse(tokens)
		result := eval.Eval(ast, env)
		fmt.Println(result)
	}
}
