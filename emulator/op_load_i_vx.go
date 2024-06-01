package emulator

import "fmt"

func (c *chip8) read(opcode Word) {
	vx := int(opcode.High() & 0xF)
	i := c.i
	for vn := 0; vn <= vx; vn++ {
		c.reg[vn] = c.mem.Read(int(i))
		i++
	}
	fmt.Printf("READ V%X\n", vx)
}
