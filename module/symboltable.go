package module

// SymbolTable holds the addresses of symbols labels (predefined, labels and variables)
type SymbolTable interface {
	Contains(symbol string) bool
	AddEntry(symbol string, address int)
	AddVariableEntry(symbol string)
	GetAddress(symbol string) (int, error)
}
