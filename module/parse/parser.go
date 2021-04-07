package parse

import (
	"fmt"
	"hack-assembler/instruction"
	"hack-assembler/utils"
	"os"
	"regexp"
	"strings"
)

const (
	comment = "//"
)

var (
	regexAInstruction = regexp.MustCompile("^@(([\\d])*|([a-zA-z_.$:][a-zA-z\\d_.$:]*)+)$")
	regexCInstruction = regexp.MustCompile("^(A?M?D?\\s*=\\s*)?([\\|&01AMD!+-]+)(\\s*[;]\\s*(JMP|JEQ|JNE|JGT|JLT|JLE|JGE)|)?$")
	regexLabel        = regexp.MustCompile("^\\([a-zA-z_.$:][a-zA-z\\d_.$:]*\\)$")
)

// Parser parses Hack assembly commands from file
type Parser struct {
	lines  []string
	cursor int
}

// NewParser returns a parser instance for a given file
func NewParser(fileName string) (*Parser, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", err.Error())
	}
	lines := utils.ReadLines(f)
	return &Parser{
		lines:  lines,
		cursor: -1,
	}, nil
}

// HasNext works as an iterator moving the cursor to the
// next line to parse and returning false what done
func (p *Parser) HasNext() bool {
	p.cursor++
	return p.cursor < len(p.lines)
}

// ParseInstruction parses the current instruction
func (p *Parser) ParseInstruction() (*instruction.Instruction, error) {
	currInstr := stripWhitespace(p.lines[p.cursor])
	// NOP instruction
	if len(currInstr) == 0 {
		return nil, nil
	}

	if regexAInstruction.MatchString(currInstr) {
		return parseAInstruction(currInstr)
	} else if regexCInstruction.MatchString(currInstr) {
		return parseCInstruction(currInstr)
	} else if regexLabel.MatchString(currInstr) {
		return parseLabel(currInstr)
	}
	return nil, fmt.Errorf("invalid instruction: %s", currInstr)
}

// Reset resets the cursor position
func (p *Parser) Reset() {
	p.cursor = -1
}

// A-instruction: @value
// Where value is either a non-negative decimal number
// or a symbol referring to such number.
func parseAInstruction(instr string) (*instruction.Instruction, error) {
	return instruction.NewAInstruction(instr[1:]), nil
}

// C-instruction: dest=comp;jump
// Either the dest or jump fields may be empty.
// If dest is empty, the ‘=’ is omitted
// If jump is empty, the ‘;’ is omitted
func parseCInstruction(instr string) (*instruction.Instruction, error) {
	var dest, comp, jmp string
	ixEq := strings.Index(instr, "=")
	ixCol := strings.Index(instr, ";")
	compStart := 0
	compEnd := len(instr)

	// At least one of dest and jmp must be present
	if ixEq == -1 && ixCol == -1 {
		return nil, fmt.Errorf("invalid instruction: %s", instr)
	}

	if ixEq != -1 {
		dest = strings.TrimSpace(instr[0:ixEq])
		compStart = ixEq + 1
	}
	if ixCol != -1 {
		jmp = strings.TrimSpace(instr[ixCol+1:])
		compEnd = ixCol
	}
	comp = strings.TrimSpace(instr[compStart:compEnd])

	return instruction.NewCInstruction(dest, comp, jmp), nil
}

// A user-defined symbol can be any sequence of letters, digits, underscore (_), dot (.),
//dollar sign ($), and colon (:) that does not begin with a digit
func parseLabel(instr string) (*instruction.Instruction, error) {
	return instruction.NewLabel(instr[1 : len(instr)-1]), nil
}

func stripWhitespace(instr string) string {
	instr = strings.Split(instr, comment)[0]
	return strings.TrimSpace(instr)
}
