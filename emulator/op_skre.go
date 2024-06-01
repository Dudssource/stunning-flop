package emulator

import "fmt"

func (c *chip8) skre(opcode Word) {
	vx := opcode.High() & 0xF
	vy := opcode.Low() & 0xF0 >> 4
	if c.reg[vx] == c.reg[vy] {
		c.pc += 2
		fmt.Printf("SKRE V%X=V%X\n", vx, vy)
	}
}
