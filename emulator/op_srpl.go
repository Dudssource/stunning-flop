package emulator

import "fmt"

func (c *chip8) srpl(opcode Word) {
	vx := opcode.High() & 0xF
	for v := uint8(0); v < vx; v++ {
		c.rpl[v] = c.reg[v]
	}
	fmt.Printf("SRPL V%X\n", vx)
}
