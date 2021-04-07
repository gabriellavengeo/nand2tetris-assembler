package instruction

// AInstruction represents an A instruction with the symbolic syntax @value
// where value is either a non-negative decimal or a symbol
type AInstruction interface {
	Value() string
}