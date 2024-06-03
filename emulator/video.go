package emulator

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Video struct {
	width          int32
	height         int32
	internalWidth  int32
	internalHeight int32
	videoMemory    [][]uint8
	scaleFactor    int32
	// SCHIP extended mode
	extended bool
}

func (v *Video) Init(width, height int32) {
	v.internalWidth = 64
	v.internalHeight = 32
	v.width = width
	v.height = height
	v.scaleFactor = width / v.internalWidth

	v.videoMemory = make([][]uint8, v.internalHeight)
	for i := range v.videoMemory {
		v.videoMemory[i] = make([]uint8, v.internalWidth)
	}

	rl.InitWindow(v.width, v.height, "Chip-8 Emulator")
	rl.SetTargetFPS(120)
}

func (v *Video) EnableExtendedMode() {
	v.extended = true
	v.internalHeight = 64
	v.internalWidth = 128
}

func (v *Video) DisableExtendedMode() {
	v.extended = false
	v.internalHeight = 32
	v.internalWidth = 64
}

func (v *Video) ScrollLeft() {
	for y := 0; y < int(v.internalHeight); y++ {
		for x := 4; x < int(v.internalWidth); x++ {
			v.videoMemory[y][x-4] = v.videoMemory[y][x]
		}
	}
	v.Draw([]byte{}, 0, 0)
}

func (v *Video) ScrollRight() {
	for y := 0; y < int(v.internalHeight); y++ {
		for x := v.internalWidth - 4; x >= 0; x-- {
			v.videoMemory[y][x+4] = v.videoMemory[y][x]
		}
	}
	v.Draw([]byte{}, 0, 0)
}

func (v *Video) Draw(bitmap []byte, x, y int32) bool {

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
				bitY = (y + int32(i)) % v.internalHeight
				bitX = (x + (7 - bit)) % v.internalWidth
			)

			// check if the k-th bit is set, by left shifting 1 by k to create a bit mask with
			// only k-th set, ANDing it with the number to extract the result and
			// right shifting by k-th again to extract the bit value (0 or 1)
			spriteBit := (sprite & (1 << bit) >> bit)

			// for SCHIP extended mode
			if v.extended {
				bitY = ((y / 2) + int32(i)) % v.internalHeight
				if i%2 == 1 {
					bitX = (x + 8 + (7 - bit)) % v.internalWidth
				}
			}

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
	for y := 0; y < int(v.internalHeight); y++ {
		for x := 0; x < int(v.internalWidth); x++ {

			var (
				posX = int32(x) * v.scaleFactor
				posY = int32(y) * v.scaleFactor
			)

			var color = rl.RayWhite

			if v.videoMemory[y][x] == 0x1 {
				color = rl.Black
			}

			rl.DrawRectangle(posX, posY, v.scaleFactor, v.scaleFactor, color)
		}
	}

	return collision
}

func (v *Video) Close() {
	rl.CloseWindow()
}

func (v *Video) Clear() {
	v.videoMemory = make([][]uint8, v.internalHeight)
	for i := range v.videoMemory {
		v.videoMemory[i] = make([]uint8, v.internalWidth)
	}
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.EndDrawing()
}
