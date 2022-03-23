package main

import (
	"fmt"
	"io"
	"lexer"
	"simple_parser"
)

func main() {
	source := "9-5+2"
	my_lexer := lexer.NewLexer(source)
	parser := simple_parser.NewSimpleParser(my_lexer)
	root, err := parser.Parse()
	if err == io.EOF {
		fmt.Println("Syntax translation: ", root.Attribute())
	} else {
		fmt.Println("source is legal expression")
	}
}
