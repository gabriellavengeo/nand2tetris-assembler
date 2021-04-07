package module

import "github.com/gabriellavengeo/nand2tetris-assembler/instruction"

// Translator translates symbolic Hack assembly instructions to binary
type Translator interface {
	TranslateAInstruction(instr instruction.AInstruction) (string, error)
	TranslateCInstruction(instr instruction.CInstruction) (string, error)
}
