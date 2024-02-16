package delphi

import (
	base "delphics/antlr"
	"fmt"
)

func (v *Visitor) EnterProgram(cx *base.ProgramContext) {
	fmt.Println("[ENTER] Program", cx.GetText())
}

func (v *Visitor) ExitProgram(cx *base.ProgramContext) {
	fmt.Println("[EXIT]  Program", cx.GetText())
	fmt.Println("[HEAD] ", cx.ProgramHead().NamespaceName().GetText())
	fmt.Println(cx.Block())
}
