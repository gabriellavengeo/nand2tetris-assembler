# nand2tetris-assembler

`nand2tetris-assembler` is a CLI tool written in Go that compiles programs written in the Hack assembly lanuage into a plain text representation of binary code for the Hack computer instruction set. The Hack assembly language and the Hack hardware platform are part of the [Nand2Tetris](https://www.nand2tetris.org/) course. 

### Build

`go build -o bin/nand2tetris-assembler ./cmd`

### Usage

```
Usage: ./nand2tetris-assembler -f [FILE] -o [FILE]
OPTIONS:
-f <string>: input file name (must have .asm extension)
-o <string>: output file name (optional)
```
