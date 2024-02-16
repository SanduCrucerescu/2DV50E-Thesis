package delphi

type Node interface {
	Enter(v AstVisitor)
}

type AstVisitor interface {
	VisitProgram(cx *Program)
	VisitVarDecl(cx *VarDecl)
	VisitPrintStatement(cx *PrintStatement)
}

// ====== File ======

type File struct {
	Section
}

type Section interface {
	Node
	EnterSection()
	VisitHead()
	VisitBody()
}

// ====== PROGRAM ======

type Program struct {
	Namespace    string
	Declarations []*VarDecl
	Statements   []Statement
}

func (p *Program) Enter(v AstVisitor) {
	v.VisitProgram(p)
}

// ====== VARIABLES ======

type VarDecl struct {
	Name string
	Type string
}

func (v *VarDecl) Enter(vis AstVisitor) {
	vis.VisitVarDecl(v)
}

// ====== STATEMENTS ======

type Statement interface {
	Node
	stmtNode()
}

type PrintStatement struct {
	Msg string
}

func (p *PrintStatement) stmtNode() {}
func (p *PrintStatement) Enter(v AstVisitor) {
	v.VisitPrintStatement(p)
}
