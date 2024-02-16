package main

import (
	"delphics/delphi"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: delphics <file.pas>")
		return
	}

	src, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	v := delphi.NewVisitor(string(src))
	tree := v.Parser.Program()

	antlr.ParseTreeWalkerDefault.Walk(v, tree)
}
