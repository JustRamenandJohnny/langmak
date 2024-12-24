package compiler

import (
	"fmt"
	"log"
	"strings"

	parser2 "langmak/pkg/parser"

	"github.com/antlr4-go/antlr/v4"
)

func NewCompiler() *Compiler {
	return &Compiler{
		heap: make(map[string]string),
		gen: &TACGenerator{
			Instructions: []TACInstruction{},
			TempCounter:  0,
			LabelCounter: 0,
		},
		code: []string{},
	}
}

type TACInstruction struct {
	Op     string
	Arg1   string
	Arg2   string
	Result string
}

type TACGenerator struct {
	Instructions []TACInstruction
	TempCounter  int
	LabelCounter int
}

func (gen *TACGenerator) NewTemp() string {
	gen.TempCounter++
	return fmt.Sprintf("t%d", gen.TempCounter)
}

func (gen *TACGenerator) NewLabel() string {
	gen.LabelCounter++
	return fmt.Sprintf("L%d", gen.LabelCounter)
}

type Compiler struct {
	heap map[string]string
	gen  *TACGenerator
	code []string
}

func (v *Compiler) GetCode() string {
	return strings.Join(v.code, "\n")
}

func (v *Compiler) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Compiler) VisitChildren(node antlr.RuleNode) interface{} {
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

func (v *Compiler) DefaultResult() interface{} {
	return nil
}

func (v *Compiler) AggregateResult(aggregate, nextResult interface{}) interface{} {
	return nextResult
}

func (v *Compiler) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (v *Compiler) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

func (v *Compiler) VisitProg(ctx *parser2.ProgContext) interface{} {
	v.VisitChildren(ctx)

	var codeStrings []string
	for _, instr := range v.gen.Instructions {
		var line string
		if instr.Arg2 != "" {
			line = fmt.Sprintf("%s = %s %s %s", instr.Result, instr.Arg1, instr.Op, instr.Arg2)
		} else if instr.Op == "label" {
			line = fmt.Sprintf("%s:", instr.Result)
		} else if instr.Op == "goto" {
			line = fmt.Sprintf("goto %s", instr.Result)
		} else if instr.Op == "ifFalse" {
			line = fmt.Sprintf("ifFalse %s goto %s", instr.Arg1, instr.Result)
		} else if instr.Op == "print" {
			line = fmt.Sprintf("print %s", instr.Arg1)
		} else if instr.Op == "=" {
			line = fmt.Sprintf("%s = %s", instr.Result, instr.Arg1)
		} else {
			line = fmt.Sprintf("%s = %s %s", instr.Result, instr.Op, instr.Arg1)
		}
		codeStrings = append(codeStrings, line)
	}

	v.code = codeStrings

	return nil
}

func (v *Compiler) VisitAssignStat(ctx *parser2.AssignStatContext) interface{} {
	varName := ctx.ID().GetText()
	value := v.Visit(ctx.Expr()).(string)
	v.heap[varName] = value

	v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
		Op:     "=",
		Arg1:   value,
		Result: varName,
	})

	return nil
}

func (v *Compiler) VisitPrintExpr(ctx *parser2.PrintExprContext) interface{} {
	value := v.Visit(ctx.Expr()).(string)

	v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
		Op:   "print",
		Arg1: value,
	})

	return nil
}

func (v *Compiler) VisitExpr(ctx *parser2.ExprContext) interface{} {
	if ctx.INT() != nil {
		return ctx.INT().GetText()
	}
	if ctx.ID() != nil {
		varName := ctx.ID().GetText()
		if _, ok := v.heap[varName]; ok {
			return varName
		} else {
			log.Fatalf("undefined variable: %s", varName)
		}
	}
	if len(ctx.AllExpr()) == 2 && ctx.GetOp() != nil {
		left := v.Visit(ctx.Expr(0)).(string)
		right := v.Visit(ctx.Expr(1)).(string)
		result := v.gen.NewTemp()

		var op string
		switch ctx.GetOp().GetTokenType() {
		case parser2.LangMakParserPLUS:
			op = "+"
		case parser2.LangMakParserMINUS:
			op = "-"
		case parser2.LangMakParserMULT:
			op = "*"
		case parser2.LangMakParserDIV:
			op = "/"
		case parser2.LangMakParserGREAT:
			op = ">"
		case parser2.LangMakParserLESS:
			op = "<"
		case parser2.LangMakParserEQUAL:
			op = "=="
		default:
			log.Fatalf("unknown operator: %s", ctx.GetOp().GetText())
		}

		v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
			Op:     op,
			Arg1:   left,
			Arg2:   right,
			Result: result,
		})

		return result
	}
	if ctx.LPAREN() != nil {
		return v.Visit(ctx.Expr(0)).(string)
	}
	log.Fatalf("unknown expression: %s", ctx.GetText())
	return nil
}

func (v *Compiler) VisitWhileStat(ctx *parser2.WhileStatContext) interface{} {
	startLabel := v.gen.NewLabel()
	endLabel := v.gen.NewLabel()

	v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
		Op:     "label",
		Result: startLabel,
	})

	condition := v.Visit(ctx.Expr()).(string)
	v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
		Op:     "ifFalse",
		Arg1:   condition,
		Result: endLabel,
	})

	for _, stat := range ctx.AllStat() {
		v.Visit(stat)
	}

	v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
		Op:     "goto",
		Result: startLabel,
	})

	v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
		Op:     "label",
		Result: endLabel,
	})

	return nil
}

func (v *Compiler) VisitIfStat(ctx *parser2.IfStatContext) interface{} {
	elseLabel := v.gen.NewLabel()
	endLabel := v.gen.NewLabel()

	condition := v.Visit(ctx.Expr()).(string)
	v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
		Op:     "ifgoto",
		Arg1:   condition,
		Result: elseLabel,
	})

	for _, stat := range ctx.AllStat() {
		v.Visit(stat)
	}

	v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
		Op:     "goto",
		Result: endLabel,
	})

	v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
		Op:     "label",
		Result: elseLabel,
	})

	v.gen.Instructions = append(v.gen.Instructions, TACInstruction{
		Op:     "label",
		Result: endLabel,
	})

	return nil
}
