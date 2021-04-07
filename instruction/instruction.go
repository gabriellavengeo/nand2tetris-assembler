package instruction

// Type represents instruction type
type Type string

const (
	// AInstructionType represents an A instruction type
	AInstructionType Type = "A"
	// CInstructionType represents an C instruction type
	CInstructionType Type = "C"
	// LabelType represents an label type
	LabelType        Type = "Label"
)

// Instruction represents a Hack assembly instruction
type Instruction struct {
	dest      string
	comp      string
	jmp       string
	value     string
	label     string
	instrType Type
}

// NewAInstruction returns a new A instruction instance
func NewAInstruction(value string) *Instruction {
	return &Instruction{
		instrType: AInstructionType,
		value:     value,
	}
}

// NewCInstruction returns a new C instruction instance
func NewCInstruction(dest, comp, jmp string) *Instruction {
	return &Instruction{
		instrType: CInstructionType,
		dest:      dest,
		comp:      comp,
		jmp:       jmp,
	}
}

// NewLabel returns a new label instance
func NewLabel(label string) *Instruction {
	return &Instruction{
		instrType: LabelType,
		label: label,
	}
}

// Type returns the instruction type (A, C or label)
func (s *Instruction) Type() Type {
	return s.instrType
}

// Value returns the instruction value for an A instruction
func (s *Instruction) Value() string {
	return s.value
}

// Dest returns the destination for a C instruction
func (s *Instruction) Dest() string {
	return s.dest
}

// Comp returns the computation for a C instruction
func (s *Instruction) Comp() string {
	return s.comp
}

// Jmp returns the jump component of a C instruction
func (s *Instruction) Jmp() string {
	return s.jmp
}

// Label returns the symbol of a label pseudo instruction
func (s *Instruction) Label() string {
	return s.label
}

