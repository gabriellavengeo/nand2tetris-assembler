package main

import (
	"flag"
	"fmt"
	"github.com/gabriellavengeo/nand2tetris-assembler/assembler/hackassembler"
	"github.com/gabriellavengeo/nand2tetris-assembler/module/parse"
	"github.com/gabriellavengeo/nand2tetris-assembler/module/symbol"
	"github.com/gabriellavengeo/nand2tetris-assembler/module/translate"

	"log"
	"os"
	"strings"
)

const (
	asmExt  = ".asm"
	hackExt = ".hack"
)

func main() {
	var asmFile, outputFile string
	flag.StringVar(&asmFile, "f", "", "input file name")
	flag.StringVar(&outputFile, "o", "", "output file name")
	flag.Parse()

	// Input file name with .asm extension must be provided
	if asmFile == "" || !strings.HasSuffix(asmFile, asmExt) {
		fmt.Println("Usage:nand2tetris-assembler -f [FILE] -o [FILE]")
		fmt.Println("OPTIONS:")
		fmt.Println("-f <string>: input file name (must have .asm extension)")
		fmt.Println("-o <string>: output file name (optional)")
		os.Exit(1)
	}

	// If no output file name is supplied use Xxx.hack as default for Xxx.asm
	if outputFile == "" {
		outputFile = strings.TrimSuffix(asmFile, asmExt) + hackExt
	}

	// Open output file, creating it if it doesn't exist and truncating it if it does
	out, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Panicf(err.Error())
	}
	defer out.Close()

	// Initialize parser, symbol table and instruction translator
	parser, err := parse.NewParser(asmFile)
	if err != nil {
		log.Panic(err.Error())
	}
	symbolTable := symbol.NewSymbolTable()
	translator := translate.NewTranslator(symbolTable)

	// Populate the symbol table and assemble and write binary code to the output file
	asm := hackassembler.NewAssembler(parser, translator, symbolTable)
	err = asm.PopulateLabelSymbols()
	if err != nil {
		log.Panic(err.Error())
	}
	err = asm.Assemble(out)
	if err != nil {
		log.Panic(err.Error())
	}
}
