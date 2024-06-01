package emulator

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Video struct {
	width       int32
	height      int32
	videoMemory [32][64]uint8
}

func (v *Video) Init(width, height int32) {
	v.width = width
	v.height = height
	rl.InitWindow(v.width, v.height, "Chip-8 Emulator")
	rl.SetTargetFPS(120)
}

func (v *Video) Draw(bitmap []byte, x, y int32) bool {

	const (
		tileSize = int32(8)
	)

	var (
		collision = false
	)

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	defer rl.EndDrawing()

	// Update
	for i, sprite := range bitmap {

		for bit := int32(7); bit >= 0; bit-- {

			var (
				// wrap pixels if they overlap the screen internal size
				bitY = (y + int32(i)) % 32
				bitX = (x + (7 - bit)) % 64
			)

			// check if the k-th bit is set, by left shifting 1 by k to create a bit mask with
			// only k-th set, ANDing it with the number to extract the result and
			// right shifting by k-th again to extract the bit value (0 or 1)
			spriteBit := (sprite & (1 << bit) >> bit)

			// check if both sprite and memory bits are set (collision), on this case set
			// the VF register with 1
			if spriteBit == 1 && v.videoMemory[bitY][bitX] == 1 {
				collision = true
			}

			// since our video memory is a 64x32 bit matrix, each bit from the N byte need to be
			// set starting from XY, for each bit read we increase X, for each N byte read we increase Y.
			v.videoMemory[bitY][bitX] ^= spriteBit
		}
	}

	// Draw
	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {

			var (
				posX = int32(x) * int32(tileSize)
				posY = int32(y) * int32(tileSize)
			)

			var color = rl.RayWhite

			if v.videoMemory[y][x] == 0x1 {
				color = rl.Black
			}

			rl.DrawRectangle(posX, posY, tileSize, tileSize, color)
		}
	}

	return collision
}

func (v *Video) Close() {
	rl.CloseWindow()
}

func (v *Video) Clear() {
	v.videoMemory = [32][64]uint8{}
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.EndDrawing()
}
