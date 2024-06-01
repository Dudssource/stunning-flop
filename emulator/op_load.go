package emulator

import "fmt"

func (c *chip8) load(opcode Word) {
	v := opcode.High() & 0xF
	nn := opcode.Low()
	c.reg[v] = nn
	fmt.Printf("LOAD V%X, $%X\n", v, nn)
}
