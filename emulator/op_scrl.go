package emulator

func (c *chip8) scrl(_ Word) {
	c.video.ScrollLeft()
}
