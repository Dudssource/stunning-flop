package emulator

import "fmt"

func (c *chip8) skp(opcode Word) {
	vx := opcode.High() & 0xF
	if c.keyboard.IsKeyDown(c.reg[vx]) {
		c.pc += 2
	}
	fmt.Printf("SKP V%X, %X\n", vx, c.reg[vx])
}
