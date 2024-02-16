package delphi

import "strings"

type Generator struct {
	Code strings.Builder
}

func (g *Generator) Write(s ...string) {
	g.Code.WriteString(strings.Join(s, ""))
}

// ========== Impl - AstVisitor ==========

func (g *Generator) VisitProgram(cx *Program) {
	g.Write("namespace ", cx.Namespace, "\n{\n")

	for _, d := range cx.Declarations {
		d.Enter(g)
	}

	g.Write("\n")
}

func (g *Generator) VisitVarDecl(cx *VarDecl) {
	g.Write("\t", cx.Type, " ", cx.Name, ";\n")
}

func (g *Generator) VisitPrintStatement(cx *PrintStatement) {
	g.Write("\t", "Console.WriteLine(\"", cx.Msg, "\");\n")
}
