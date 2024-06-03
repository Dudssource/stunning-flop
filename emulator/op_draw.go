package emulator

func (c *chip8) draw(opcode Word) {
	vx := opcode.High() & 0xF
	vy := opcode.Low() & 0xF0 >> 4
	n := int(opcode.Low() & 0xF)
	rowWiseBitmap := false

	// SCHIP 16x16 sprite instruction
	if n == 0 && c.video.extended {
		n = 32
		rowWiseBitmap = true
	}

	if n > 0 {
		bitmap := make([]byte, n)
		temp := c.i
		for i := 0; i < n; i++ {
			sprite := c.mem.Read(int(temp))
			temp++
			bitmap[i] = sprite
		}

		if c.video.Draw(bitmap, int32(c.reg[vx]), int32(c.reg[vy]), rowWiseBitmap) {
			c.reg[0xF] = 1
		}
	}
}
