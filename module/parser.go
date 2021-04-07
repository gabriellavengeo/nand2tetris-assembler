package module

import "hack-assembler/instruction"

// Parser parses Hack assembly commands from file
type Parser interface {
	HasNext() bool
	ParseInstruction() (*instruction.Instruction, error)
	Reset()
}
