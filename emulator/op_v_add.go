package emulator

import "fmt"

func (c *chip8) addr(opcode Word) {
	vx := opcode.High() & 0xF
	vy := opcode.Low() & 0xF0 >> 4
	temp := uint16(c.reg[vx] + c.reg[vy])
	if temp > 0xFF {
		c.reg[0xF] = 1
		c.reg[vx] = uint8(temp - 256)
	} else {
		c.reg[0xF] = 0
		c.reg[vx] = uint8(temp)
	}
	fmt.Printf("ADD V%x, V%x, VF=%x\n", vx, vy, c.reg[0xF])
}
