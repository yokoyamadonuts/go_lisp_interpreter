package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Lisp REPL. Type 'exit' to quit.")
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			break
		}
		fmt.Println("You entered:", text)
	}
}
