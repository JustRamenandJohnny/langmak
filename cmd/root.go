package cmd

import (
	"log"
	"os"

	"langmak/pkg/compiler"
	"langmak/pkg/interpreter"
	parser2 "langmak/pkg/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string

func Run(inputFile, outputFile string) {
	if inputFile == "" {
		log.Fatal("Error: --input flag is required. Please specify a LangMak file.")
	}

	input, err := antlr.NewFileStream(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	lexer := parser2.NewLangMakLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser2.NewLangMakParser(stream)
	tree := p.Prog()

	if outputFile != "" {
		compiler := compiler.NewCompiler()
		compiler.Visit(tree)
		compiledCode := compiler.GetCode()

		err := saveToFile(outputFile, compiledCode)
		if err != nil {
			log.Fatalf("Error saving to file: %s", err)
		}

	} else {
		interpreter := interpreter.NewInterpreter()
		interpreter.Visit(tree)
	}
}

var RootCmd = &cobra.Command{
	Use:   "micro-lang",
	Short: "LangMak - a language for micromashinki/noVax and VAX compatibility",
	Long: `LangMak is a programming language designed for micromashinki/noVax 
and fully compatible with VAX systems. Use this tool to interpret or compile LangMak programs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Run(inputFile, outputFile)
	},
}

func saveToFile(path string, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

func init() {
	RootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to the input LangMak file (required)")
	RootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Path to the output file (optional)")
	RootCmd.MarkFlagRequired("input")
}
