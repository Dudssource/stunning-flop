package emulator

import "fmt"

func (c *chip8) rrpl(opcode Word) {
	vx := opcode.High() & 0xF
	for v := uint8(0); v < vx; v++ {
		c.reg[v] = c.rpl[v]
	}
	fmt.Printf("RRPL V%X\n", vx)
}
