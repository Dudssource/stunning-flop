package emulator

func (c *chip8) scrr(_ Word) {
	c.video.ScrollRight()
}
