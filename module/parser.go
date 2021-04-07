package module

import "github.com/gabriellavengeo/nand2tetris-assembler/instruction"

// Parser parses Hack assembly commands from file
type Parser interface {
	HasNext() bool
	ParseInstruction() (*instruction.Instruction, error)
	Reset()
}
