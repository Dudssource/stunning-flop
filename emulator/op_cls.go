package emulator

func (c *chip8) cls(_ Word) {
	c.video.Clear()
}
