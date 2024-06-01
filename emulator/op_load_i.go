package emulator

import "fmt"

func (c *chip8) load_i(opcode Word) {
	addr := Word(opcode.Word() & 0x0FFF)
	c.i = addr
	fmt.Printf("LOAD I, %03X\n", addr)
}
