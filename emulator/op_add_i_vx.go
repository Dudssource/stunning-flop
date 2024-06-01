package emulator

import "fmt"

func (c *chip8) add_i(opcode Word) {
	vx := opcode.High() & 0xF
	c.i += Word(c.reg[vx])
	fmt.Printf("ADDI V%X\n", vx)
}
