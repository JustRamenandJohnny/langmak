// Code generated from LangMak.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // LangMak

import "github.com/antlr4-go/antlr/v4"

// LangMakListener is a complete listener for a parse tree produced by LangMakParser.
type LangMakListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterAssignStat is called when entering the assignStat production.
	EnterAssignStat(c *AssignStatContext)

	// EnterPrintExpr is called when entering the printExpr production.
	EnterPrintExpr(c *PrintExprContext)

	// EnterWhileStat is called when entering the whileStat production.
	EnterWhileStat(c *WhileStatContext)

	// EnterIfStat is called when entering the ifStat production.
	EnterIfStat(c *IfStatContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitAssignStat is called when exiting the assignStat production.
	ExitAssignStat(c *AssignStatContext)

	// ExitPrintExpr is called when exiting the printExpr production.
	ExitPrintExpr(c *PrintExprContext)

	// ExitWhileStat is called when exiting the whileStat production.
	ExitWhileStat(c *WhileStatContext)

	// ExitIfStat is called when exiting the ifStat production.
	ExitIfStat(c *IfStatContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
