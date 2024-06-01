package emulator

import "fmt"

func (c *chip8) sub(opcode Word) {
	vx := opcode.High() & 0xF
	vy := opcode.Low() & 0xF0 >> 4
	if c.reg[vx] > c.reg[vy] {
		c.reg[0xF] = 1
		c.reg[vx] -= c.reg[vy]
	} else {
		c.reg[0xF] = 0
		c.reg[vx] = uint8(256 + int(c.reg[vx]-c.reg[vy]))
	}
	fmt.Printf("SUB V%X, V%X, VF=%x\n", vx, vy, c.reg[0xF])
}
