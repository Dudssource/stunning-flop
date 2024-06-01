package emulator

import "fmt"

func (c *chip8) skne(opcode Word) {
	v := opcode.High() & 0xF
	if c.reg[v] != opcode.Low() {
		c.pc += 2
		fmt.Printf("SKNE V%X!=%X\n", v, opcode.Low())
	}
}
