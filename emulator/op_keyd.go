package emulator

import "fmt"

func (c *chip8) keyd(opcode Word) {
	vx := opcode.High() & 0xF
	if pin, ok := c.keyboard.WaitForKeyPressed(); ok {
		c.reg[vx] = pin
		fmt.Printf("KEYD V%X, %X\n", vx, c.reg[vx])
	} else {
		// go back PC in two
		c.pc -= 2
	}
}
