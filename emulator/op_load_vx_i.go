package emulator

import "fmt"

func (c *chip8) stor(opcode Word) {
	vx := int(opcode.High() & 0xF)
	i := c.i
	for vn := 0; vn <= vx; vn++ {
		c.mem.Write(i, c.reg[vn])
		i++
	}
	fmt.Printf("STOR V%X\n", vx)
}
