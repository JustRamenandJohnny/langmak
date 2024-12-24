// Code generated from LangMak.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // LangMak

import "github.com/antlr4-go/antlr/v4"

// BaseLangMakListener is a complete listener for a parse tree produced by LangMakParser.
type BaseLangMakListener struct{}

var _ LangMakListener = &BaseLangMakListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLangMakListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLangMakListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLangMakListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLangMakListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseLangMakListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseLangMakListener) ExitProg(ctx *ProgContext) {}

// EnterAssignStat is called when production assignStat is entered.
func (s *BaseLangMakListener) EnterAssignStat(ctx *AssignStatContext) {}

// ExitAssignStat is called when production assignStat is exited.
func (s *BaseLangMakListener) ExitAssignStat(ctx *AssignStatContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *BaseLangMakListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *BaseLangMakListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterWhileStat is called when production whileStat is entered.
func (s *BaseLangMakListener) EnterWhileStat(ctx *WhileStatContext) {}

// ExitWhileStat is called when production whileStat is exited.
func (s *BaseLangMakListener) ExitWhileStat(ctx *WhileStatContext) {}

// EnterIfStat is called when production ifStat is entered.
func (s *BaseLangMakListener) EnterIfStat(ctx *IfStatContext) {}

// ExitIfStat is called when production ifStat is exited.
func (s *BaseLangMakListener) ExitIfStat(ctx *IfStatContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseLangMakListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseLangMakListener) ExitExpr(ctx *ExprContext) {}
