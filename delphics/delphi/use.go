package delphi

import (
	base "delphics/antlr"
	"fmt"
)

func (v *Visitor) EnterUsesClause(cx *base.UsesClauseContext) {
	fmt.Println("[ENTER] Uses", cx.GetText())
}

func (v *Visitor) ExitUsesClause(cx *base.UsesClauseContext) {
	fmt.Println("[EXIT] Uses", cx.GetText())

	v.Uses = append(v.Uses, cx.GetText())
}
