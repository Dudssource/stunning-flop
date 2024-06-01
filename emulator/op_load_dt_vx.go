package emulator

import "fmt"

func (c *chip8) move_d(opcode Word) {
	vx := opcode.High() & 0xF
	c.reg[vx] = c.dt
	fmt.Printf("MOVED V%X, DT\n", vx)
}
