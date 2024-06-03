package emulator

func (c *chip8) scrd(opcode Word) {
	n := opcode.Low() & 0xF
	c.video.ScrollDown(n)
}
