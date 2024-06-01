package emulator

import (
	"fmt"
	"math/rand"
)

func (c *chip8) rnd(opcode Word) {
	vx := opcode.High() & 0xF
	kk := opcode.Low()
	c.reg[vx] = uint8(rand.Intn(0xFF)) & kk
	fmt.Printf("RND V%X, %X\n", vx, kk)
}
