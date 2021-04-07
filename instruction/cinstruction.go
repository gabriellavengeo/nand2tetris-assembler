package instruction

// CInstruction represents an C instruction with the symbolic syntax dest = comp ; jmp
type CInstruction interface {
	Dest() string
	Comp() string
	Jmp() string
}
