package interpreter

import (
	"os"
	"testing"

	"langmak/pkg/parser"

	"github.com/antlr4-go/antlr/v4"
)

func TestInterpreter(t *testing.T) {
	tests := map[string]int{
		"testdata/test1.micro": 7,
		"testdata/test2.micro": 1,
		"testdata/test3.micro": 9,
	}

	for path, expected := range tests {
		input, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("Can not read file %s: %v", path, err)
		}

		inputStream := antlr.NewInputStream(string(input))
		lexer := parser.NewLangMakLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewLangMakParser(tokenStream)

		tree := p.Prog()

		interp := NewInterpreter()
		interp.Visit(tree)

		actual, ok := interp.variables["result"]
		if !ok {
			t.Errorf("'result' not defined: %s", path)
			continue
		}

		if actual != expected {
			t.Errorf("Incorrect value %s %d %d", path, expected, actual)
		}
	}
}
