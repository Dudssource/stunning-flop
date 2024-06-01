package emulator

import "fmt"

func (c *chip8) shl(opcode Word) {
	vx := opcode.High() & 0xF
	if (c.reg[vx]&0x80)>>7 == 1 {
		c.reg[0xF] = 1
	} else {
		c.reg[0xF] = 0
	}
	c.reg[vx] <<= 1
	fmt.Printf("SHL V%X, VF=%X\n", vx, c.reg[0xF])
}
