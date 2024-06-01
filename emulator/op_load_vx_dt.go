package emulator

import "fmt"

func (c *chip8) load_d(opcode Word) {
	v := opcode.High() & 0xF
	c.dt = c.reg[v]
	fmt.Printf("LOADD, V%X\n", v)
}
