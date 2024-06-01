package emulator

import "fmt"

func (c *chip8) bcd(opcode Word) {
	vx := opcode.High() & 0xF
	hundreds := c.reg[vx] / 100
	tens := (c.reg[vx] % 100) / 10
	ones := (c.reg[vx] % 100) % 10
	c.mem.Write(c.i, hundreds)
	c.mem.Write(c.i+1, tens)
	c.mem.Write(c.i+2, ones)
	fmt.Printf("BCD V%X\n", vx)
}
