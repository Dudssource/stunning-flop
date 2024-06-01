package emulator

import "fmt"

func (c *chip8) move(opcode Word) {
	vx := opcode.High() & 0xF
	vy := opcode.Low() & 0xF0 >> 4
	c.reg[vx] = c.reg[vy]
	fmt.Printf("MOVE V%X, V%X\n", vx, vy)
}
