package emulator

import "fmt"

func (c *chip8) call(opcode Word) {
	c.mem.Write(c.sp, c.pc.Low())
	c.sp++
	c.mem.Write(c.sp, c.pc.High())
	c.sp++
	c.pc = opcode & 0x0FFF
	fmt.Printf("CALL %x\n", c.pc)
}
