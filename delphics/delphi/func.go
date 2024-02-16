package delphi

import (
	base "delphics/antlr"
	"fmt"
)

func (v *Visitor) EnterFunctionDirective(cx *base.FunctionDirectiveContext) {
	fmt.Println("[ENTER] Function", cx.GetText())
}

func (v *Visitor) ExitFunctionDirective(cx *base.FunctionDirectiveContext) {
	fmt.Println("[EXIT]  Function", cx.GetText())
}
