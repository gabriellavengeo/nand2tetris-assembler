package symbol

import (
	"fmt"
)

// SymbolTable holds the addresses of symbols labels (predefined, labels and variables)
type SymbolTable struct {
	symbols        map[string]int
	nextRAMAddress int
}

// NewSymbolTable returns a new symbol table instance and populates it with the predefines symbols
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols:        predefinedSymbols,
		nextRAMAddress: 16,
	}
}

// Contains checks if the symbol table contain the given symbol
func (s *SymbolTable) Contains(symbol string) bool {
	_, ok := s.symbols[symbol]
	return ok
}

// AddEntry adds the pair (symbol, address) to the table
func (s *SymbolTable) AddEntry(symbol string, address int) {
	s.symbols[symbol] = address
}

// AddVariableEntry adds a symbol to the table with the next available RAM address
func (s *SymbolTable) AddVariableEntry(symbol string) {
	s.AddEntry(symbol, s.nextRAMAddress)
	s.nextRAMAddress += 1
}

// GetAddress returns the address associated with the symbol
func (s *SymbolTable) GetAddress(symbol string) (int, error) {
	if !s.Contains(symbol) {
		return 0, fmt.Errorf("could not resolve symbol: %s", symbol)
	}
	return s.symbols[symbol], nil
}
