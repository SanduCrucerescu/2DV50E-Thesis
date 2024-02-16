package delphi

import "github.com/antlr4-go/antlr/v4"

type FnSyntaxError func(errors chan SynErr, line, column int, msg string, e antlr.RecognitionException)

type Src struct {
	Start, End, Line, Col int
	Text                  string
}
