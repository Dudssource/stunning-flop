package emulator

import rl "github.com/gen2brain/raylib-go/raylib"

var pinpad = map[uint8]int32{
	0x0: rl.KeyX,
	0x1: rl.KeyOne,
	0x2: rl.KeyTwo,
	0x3: rl.KeyThree,
	0x4: rl.KeyQ,
	0x5: rl.KeyW,
	0x6: rl.KeyE,
	0x7: rl.KeyA,
	0x8: rl.KeyS,
	0x9: rl.KeyD,
	0xA: rl.KeyZ,
	0xB: rl.KeyC,
	0xC: rl.KeyFour,
	0xD: rl.KeyR,
	0xE: rl.KeyF,
	0xF: rl.KeyV,
}

var keys = map[int32]uint8{
	rl.KeyX:     0x0,
	rl.KeyOne:   0x1,
	rl.KeyTwo:   0x2,
	rl.KeyThree: 0x3,
	rl.KeyQ:     0x4,
	rl.KeyW:     0x5,
	rl.KeyE:     0x6,
	rl.KeyA:     0x7,
	rl.KeyS:     0x8,
	rl.KeyD:     0x9,
	rl.KeyZ:     0xA,
	rl.KeyC:     0xB,
	rl.KeyFour:  0xC,
	rl.KeyR:     0xD,
	rl.KeyF:     0xE,
	rl.KeyV:     0xF,
}

type Keyboard struct {
}

func (*Keyboard) IsKeyDown(pin uint8) bool {
	return rl.IsKeyPressed(pinpad[pin]) || rl.IsKeyDown(pinpad[pin])
}

func (*Keyboard) IsKeyUp(key uint8) bool {
	return rl.IsKeyUp(pinpad[key])
}

func (*Keyboard) WaitForKeyPressed() (uint8, bool) {

	for {
		key := rl.GetKeyPressed()
		if key == 0 {
			return 0, false
		}

		if pin, ok := keys[key]; ok {
			return pin, true
		} else {
			return 0, false
		}
	}
}
