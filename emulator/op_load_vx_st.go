package emulator

import "fmt"

func (c *chip8) loads(opcode Word) {
	v := opcode.High() & 0xF
	c.st = c.reg[v]
	fmt.Printf("LOADS, V%X\n", v)
}
