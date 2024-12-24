#!/usr/bin/make -f
SHELL = bash

BIN=bin
ANTLR_BIN=$(BIN)/antlr-4.13.2-complete.jar

GRAMMAR_FILE=LangMak.g4
GRAMMAR_PATH=$(GRAMMAR_FILE)
GRAMMAR_OUTPUT=pkg/parser

generate:
	@echo "⇒ Generate parser"
	java -jar $(ANTLR_BIN)  -visitor -Dlanguage=Go -o $(GRAMMAR_OUTPUT) $(GRAMMAR_PATH)

test: GOFLAGS ?= "-count=1"
test:
	@echo "⇒ Running go test"
	@GOFLAGS="$(GOFLAGS)" go test ./...