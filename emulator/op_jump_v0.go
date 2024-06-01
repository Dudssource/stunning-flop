package emulator

import "fmt"

func (c *chip8) jump_i(opcode Word) {
	addr := Word(opcode & 0x0FFF)
	c.pc = addr + c.i
	fmt.Printf("JUMPI, %X\n", c.pc)
}
