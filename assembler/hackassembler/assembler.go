package hackassembler

import (
	"fmt"
	"github.com/gabriellavengeo/nand2tetris-assembler/instruction"
	"github.com/gabriellavengeo/nand2tetris-assembler/module"
	"io"
	"log"
)

// Assembler converts symbolic Hack assembly code to binary
type Assembler struct {
	parser      module.Parser
	translator  module.Translator
	symbolTable module.SymbolTable
}

// NewAssembler returns a new Assembler instance
func NewAssembler(parser module.Parser, translator module.Translator,
	symbolTable module.SymbolTable) *Assembler {
	return &Assembler{
		parser:      parser,
		translator:  translator,
		symbolTable: symbolTable,
	}
}

// PopulateLabelSymbols adds entries in the symbol table for the user defined labels
func (a *Assembler) PopulateLabelSymbols() error {
	// Reset the cursor of the parser
	a.parser.Reset()

	// Iterate over the instructions and add entries in the symbol table for the labels
	var addr int
	for a.parser.HasNext() {
		instr, err := a.parser.ParseInstruction()
		if err != nil {
			return err
		}
		if instr == nil {
			continue
		}

		switch instr.Type() {
		case instruction.AInstructionType:
			addr++
		case instruction.CInstructionType:
			addr++
		case instruction.LabelType:
			a.symbolTable.AddEntry(instr.Label(), addr)
		default:
			return fmt.Errorf("unknown instruction type: %s", instr.Type())
		}
	}
	return nil
}

// Assemble converts they symbolic assembly code from a file to binary and writes to an io.Writer
func (a *Assembler) Assemble(w io.Writer) error {
	// Reset the cursor of the parser
	a.parser.Reset()

	// Iterate over the instructions, parse them and translate them to binary code
	for a.parser.HasNext() {
		instr, err := a.parser.ParseInstruction()
		if err != nil {
			return err
		}
		if instr == nil {
			continue
		}

		var instrBinary string
		switch instr.Type() {
		case instruction.AInstructionType:
			instrBinary, err = a.translator.TranslateAInstruction(instr)
			if err != nil {
				log.Panic(err.Error())
			}
		case instruction.CInstructionType:
			instrBinary, err = a.translator.TranslateCInstruction(instr)
			if err != nil {
				return err
			}
		case instruction.LabelType:
			// Skip label pseudo instructions
			continue
		default:
			return fmt.Errorf("unknown instruction type: %s", instr.Type())
		}

		_, err = io.WriteString(w, fmt.Sprintf("%s\n", instrBinary))
		if err != nil {
			return err
		}
	}
	return nil
}
