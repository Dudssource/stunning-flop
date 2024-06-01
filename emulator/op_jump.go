package emulator

import "fmt"

func (c *chip8) jump(opcode Word) {
	c.pc = opcode & 0x0FFF
	fmt.Printf("JUMP %x\n", c.pc)
}
