package translate

import (
	"fmt"
	"github.com/gabriellavengeo/nand2tetris-assembler/instruction"
	"github.com/gabriellavengeo/nand2tetris-assembler/module/symbol"

	"strconv"
)

// Translator translates symbolic Hack assembly instructions to binary
type Translator struct {
	symbolTable *symbol.SymbolTable
}

// NewTranslator returns a new Translator instance
func NewTranslator(symbolTable *symbol.SymbolTable) *Translator {
	return &Translator{symbolTable: symbolTable}
}

// TranslateAInstruction translates an A instruction to binary
func (t *Translator) TranslateAInstruction(instr instruction.AInstruction) (string, error) {
	value, err := t.translateValue(instr.Value())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0%s", value), nil
}

// TranslateCInstruction translates a C instruction to binary
func (t *Translator) TranslateCInstruction(instr instruction.CInstruction) (string, error) {
	comp, err := t.translateComp(instr.Comp())
	if err != nil {
		return "", err
	}
	dest, err := t.translateDest(instr.Dest())
	if err != nil {
		return "", err
	}
	jmp, err := t.translateJmp(instr.Jmp())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("111%s%s%s", comp, dest, jmp), nil
}

func (t *Translator) translateDest(dest string) (string, error) {
	val, ok := destMnemonics[dest]
	if !ok {
		return "", fmt.Errorf("invalid dest: %s", dest)
	}
	return val, nil
}

func (t *Translator) translateJmp(jmp string) (string, error) {
	val, ok := jmpMnemonics[jmp]
	if !ok {
		return "", fmt.Errorf("invalid jmp: %s", jmp)
	}
	return val, nil
}

func (t *Translator) translateComp(comp string) (string, error) {
	val, ok := compMnemonics[comp]
	if !ok {
		return "", fmt.Errorf("invalid comp: %s", comp)
	}
	return val, nil
}

func (t *Translator) translateValue(value string) (string, error) {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		// If the value is not an integer look up symbol table
		// In case symbol is not present, add symbol fo a new variable
		if !t.symbolTable.Contains(value) {
			t.symbolTable.AddVariableEntry(value)
		}
		intVal, err = t.symbolTable.GetAddress(value)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%.15b", intVal), nil
}
