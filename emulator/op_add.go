package emulator

import "fmt"

func (c *chip8) add(opcode Word) {
	vx := opcode.High() & 0xF
	nn := opcode.Low()
	temp := uint16(c.reg[vx] + nn)
	if temp > 0xFF {
		c.reg[0xF] = 0x1
		c.reg[vx] = uint8(temp - 256)
	} else {
		c.reg[0xF] = 0x0
		c.reg[vx] = uint8(temp)
	}
	fmt.Printf("ADD V%X, %X\n", vx, nn)
}
