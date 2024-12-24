package interpreter

import (
	"fmt"
	"log"
	"strconv"

	parser2 "langmak/pkg/parser"

	"github.com/antlr4-go/antlr/v4"
)

func NewInterpreter() *Interpreter {
	return &Interpreter{
		variables: make(map[string]int),
	}
}

type Interpreter struct {
	variables map[string]int
}

func (v *Interpreter) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Interpreter) VisitChildren(node antlr.RuleNode) interface{} {
	var result = v.DefaultResult()
	for i := 0; i < node.GetChildCount(); i++ {
		c := node.GetChild(i)
		if child, ok := c.(antlr.ParseTree); ok {
			childResult := child.Accept(v)
			result = v.AggregateResult(result, childResult)
		}
	}
	return result
}

func (v *Interpreter) DefaultResult() interface{} {
	return nil
}

func (v *Interpreter) AggregateResult(aggregate, nextResult interface{}) interface{} {
	return nextResult
}

func (v *Interpreter) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (v *Interpreter) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

func (v *Interpreter) VisitProg(ctx *parser2.ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Interpreter) VisitAssignStat(ctx *parser2.AssignStatContext) interface{} {
	varName := ctx.ID().GetText()
	value := v.Visit(ctx.Expr()).(int)
	v.variables[varName] = value
	return nil
}

func (v *Interpreter) VisitPrintExpr(ctx *parser2.PrintExprContext) interface{} {
	value := v.Visit(ctx.Expr()).(int)
	fmt.Println(value)
	return nil
}

func (v *Interpreter) VisitExpr(ctx *parser2.ExprContext) interface{} {
	if ctx.INT() != nil {
		value, err := strconv.Atoi(ctx.INT().GetText())
		if err != nil {
			panic(err)
		}
		return value
	}
	if ctx.ID() != nil {
		varName := ctx.ID().GetText()
		if value, ok := v.variables[varName]; ok {
			return value
		} else {
			log.Fatalf("undefined variable: %s", varName)
		}
	}
	if len(ctx.AllExpr()) == 2 && ctx.GetOp() != nil {
		left := v.Visit(ctx.Expr(0)).(int)
		right := v.Visit(ctx.Expr(1)).(int)
		var result int
		switch ctx.GetOp().GetTokenType() {
		case parser2.LangMakParserPLUS:
			result = left + right
		case parser2.LangMakParserMINUS:
			result = left - right
		case parser2.LangMakParserMULT:
			result = left * right
		case parser2.LangMakParserDIV:
			result = left / right
		case parser2.LangMakParserGREAT:
			if left > right {
				result = 1
			} else {
				result = 0
			}
		case parser2.LangMakParserLESS:
			if left < right {
				result = 1
			} else {
				result = 0
			}
		case parser2.LangMakParserEQUAL:
			if left == right {
				result = 1
			} else {
				result = 0
			}
		default:
			log.Fatalf("unknown operator: %s", ctx.GetOp().GetText())
		}
		return result
	}
	if ctx.LPAREN() != nil {
		return v.Visit(ctx.Expr(0)).(int)
	}
	log.Fatalf("unknown expression: %s", ctx.GetText())
	return nil
}

func (v *Interpreter) VisitWhileStat(ctx *parser2.WhileStatContext) interface{} {
	for v.Visit(ctx.Expr()).(int) == 1 {
		for _, stat := range ctx.AllStat() {
			v.Visit(stat)
		}
	}
	return nil
}

func (v *Interpreter) VisitIfStat(ctx *parser2.IfStatContext) interface{} {
	if v.Visit(ctx.Expr()).(int) == 1 {
		for _, stat := range ctx.AllStat() {
			v.Visit(stat)
		}
	}
	return nil
}
