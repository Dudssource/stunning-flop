package emulator

import "fmt"

func (c *chip8) ldspr(opcode Word) {
	vx := opcode.High() & 0xF
	c.i = Word(c.reg[vx] * 5)
	fmt.Printf("LDSPR I, V%X\n", vx)
}
