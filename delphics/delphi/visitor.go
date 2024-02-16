package delphi

import (
	base "delphics/antlr"
	"fmt"
	"log"
	"os"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	base.BaseDelphiListener

	syms  []string
	rules []string

	Prog Program

	Lexer  *base.DelphiLexer
	Parser *base.DelphiParser
	Stream *antlr.InputStream

	Errors       chan SynErr
	OnSyntaxErr  FnSyntaxError
	OnDefaultErr func()
}

// ============== BEGIN - Antlr Error Listener Implementation ==============

func (v *Visitor) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	v.OnSyntaxErr(v.Errors, line, column, msg, e)
}

func DefaultOnSyntaxErr(errors chan SynErr, line, column int, msg string, e antlr.RecognitionException) {
	log.Println("SyntaxError")
	errors <- SynErr{
		Line:      line,
		Col:       column,
		Msg:       msg,
		Exception: e,
	}
	log.Println(<-errors)
	os.Exit(1)
}

func (v *Visitor) ReportAmbiguity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex, stopIndex int,
	exact bool,
	ambigAlts *antlr.BitSet,
	configs *antlr.ATNConfigSet,
) {
	log.Println("ReportAmbiguity")
}

func (v *Visitor) ReportAttemptingFullContext(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex, stopIndex int,
	conflictingAlts *antlr.BitSet,
	configs *antlr.ATNConfigSet,
) {
	log.Println("ReportAttemptingFullContext")
	fmt.Println("  " + recognizer.GetCurrentToken().GetText())
}

func (v *Visitor) ReportContextSensitivity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex, stopIndex, prediction int,
	configs *antlr.ATNConfigSet,
) {
	log.Println("ReportContextSensitivity")
}

// ============== END - Antlr Error Listener Implementation ==============

func NewVisitor(src string) *Visitor {
	stream := antlr.NewInputStream(src)
	lexer := base.NewDelphiLexer(stream)
	tokStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := base.NewDelphiParser(tokStream)
	errs := make(chan SynErr, 1)
	onErr := func() {
		e := <-errs
		fmt.Printf("[ERROR] %d:%d\n%s\n%*s^\n%s\n",
			e.Line, e.Col,
			strings.Split(src, "\n")[e.Line-1][:e.Col+1],
			e.Col, "^",
			e.Msg,
		)
		os.Exit(1)
	}

	syms, rules := lexer.GetSymbolicNames(), parser.GetRuleNames()
	vis := &Visitor{
		syms:  syms,
		rules: rules,

		Prog: Program{
			Declarations: []*VarDecl{},
			Statements:   []Statement{},
		},

		Lexer:        lexer,
		Parser:       parser,
		Stream:       stream,
		Errors:       errs,
		OnDefaultErr: onErr,
		OnSyntaxErr:  DefaultOnSyntaxErr,
	}

	parser.RemoveErrorListeners()
	parser.AddErrorListener(vis)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(vis)

	return vis
}
