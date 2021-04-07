package symbol

import "fmt"

var predefinedSymbols = map[string]int{
	"SP":     0,     // 0x0000
	"LCL":    1,     // 0x0001
	"ARG":    2,     // 0x0002
	"THIS":   3,     // 0x0003
	"THAT":   4,     // 0x0004
	"SCREEN": 16384, // 0x4000
	"KBD":    24576, // 0x6000,
}

func init() {
	//R0-R15 0-15 0x0000-f
	for i := 0; i <= 15; i++ {
		predefinedSymbols[fmt.Sprintf("R%d", i)] = i
	}
}
