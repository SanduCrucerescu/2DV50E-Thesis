package main

import (
	"delphics/delphi"
	"log"
	"os"

	antlr "github.com/antlr4-go/antlr/v4"
)

func main() {
	src, err := os.ReadFile("data/HelloWorld.pas")
	if err != nil {
		log.Fatal(err)
		return
	}

	v := delphi.NewVisitor(string(src))
	tree := v.Parser.Program()

	antlr.ParseTreeWalkerDefault.Walk(v, tree)
}
