package emulator

import "fmt"

func (c *chip8) ske(opcode Word) {
	v := opcode.High() & 0xF
	if c.reg[v] == opcode.Low() {
		c.pc += 2
		fmt.Printf("SKE v%X=%X\n", v, opcode.Low())
	}
}
