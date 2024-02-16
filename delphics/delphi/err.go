package delphi

import (
	antlr "github.com/antlr4-go/antlr/v4"
)

type SynErr struct {
	Line, Col int
	Msg       string
	Exception antlr.RecognitionException
}
