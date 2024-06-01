package emulator

import "fmt"

func (c *chip8) rts() {
	c.sp--
	hi := uint16(c.mem.Read(int(c.sp)))
	c.sp--
	lo := uint16(c.mem.Read(int(c.sp)))
	c.pc = Word((hi << 8) | (lo & 0xFF))
	fmt.Printf("RTS %X\n", c.pc)
}
