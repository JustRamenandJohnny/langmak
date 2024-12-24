// Code generated from LangMak.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // LangMak

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by LangMakParser.
type LangMakVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LangMakParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by LangMakParser#assignStat.
	VisitAssignStat(ctx *AssignStatContext) interface{}

	// Visit a parse tree produced by LangMakParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by LangMakParser#whileStat.
	VisitWhileStat(ctx *WhileStatContext) interface{}

	// Visit a parse tree produced by LangMakParser#ifStat.
	VisitIfStat(ctx *IfStatContext) interface{}

	// Visit a parse tree produced by LangMakParser#expr.
	VisitExpr(ctx *ExprContext) interface{}
}
