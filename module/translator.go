package module

import "hack-assembler/instruction"

// Translator translates symbolic Hack assembly instructions to binary
type Translator interface {
	TranslateAInstruction(instr instruction.AInstruction) (string, error)
	TranslateCInstruction(instr instruction.CInstruction) (string, error)
}
