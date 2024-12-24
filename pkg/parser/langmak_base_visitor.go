// Code generated from LangMak.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // LangMak

import "github.com/antlr4-go/antlr/v4"

type BaseLangMakVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLangMakVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLangMakVisitor) VisitAssignStat(ctx *AssignStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLangMakVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLangMakVisitor) VisitWhileStat(ctx *WhileStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLangMakVisitor) VisitIfStat(ctx *IfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLangMakVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}
