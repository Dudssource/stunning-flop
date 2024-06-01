package emulator

import "fmt"

func (c *chip8) sknp(opcode Word) {
	vx := opcode.High() & 0xF
	if !c.keyboard.IsKeyDown(c.reg[vx]) {
		c.pc += 2
	}
	fmt.Printf("SKNP V%X, %X\n", vx, c.reg[vx])
}
