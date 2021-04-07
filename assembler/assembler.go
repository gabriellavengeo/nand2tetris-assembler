package assembler

import (
	"io"
)

// Assembler converts symbolic assembly code to binary
type Assembler interface {
	PopulateLabelSymbols() error
	Assemble(w io.Writer) error
}
