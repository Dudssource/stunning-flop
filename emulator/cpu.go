package emulator

import (
	"fmt"
	"io"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var fonts = []byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0,
	0x20, 0x60, 0x20, 0x20, 0x70,
	0xF0, 0x10, 0xF0, 0x80, 0xF0,
	0xF0, 0x10, 0xF0, 0x10, 0xF0,
	0x90, 0x90, 0xF0, 0x10, 0x10,
	0xF0, 0x80, 0xF0, 0x10, 0xF0,
	0xF0, 0x80, 0xF0, 0x90, 0xF0,
	0xF0, 0x10, 0x20, 0x40, 0x40,
	0xF0, 0x90, 0xF0, 0x90, 0xF0,
	0xF0, 0x90, 0xF0, 0x10, 0xF0,
	0xF0, 0x90, 0xF0, 0x90, 0x90,
	0xE0, 0x90, 0xE0, 0x90, 0xE0,
	0xF0, 0x80, 0x80, 0x80, 0xF0,
	0xE0, 0x90, 0x90, 0x90, 0xE0,
	0xF0, 0x80, 0xF0, 0x80, 0xF0,
	0xF0, 0x80, 0xF0, 0x80, 0x80,
}

type chip8 struct {

	// reg General Purpose 8-bit registers
	reg []uint8

	// i Index Register
	i Word

	// st Sound timer
	st byte

	// dt Delay timer
	dt byte

	// pc Stack Pointer
	sp Word

	// pc Program Counter
	pc Word

	// pc 4kb memory
	mem *Memory

	// video
	video *Video

	// Keyboard device
	keyboard *Keyboard

	// SCHIP - cpu should stop
	stop bool

	// SCHIP rpl flag registers
	rpl []uint8
}

func Cpu() *chip8 {
	return &chip8{
		reg:      make([]byte, 0x10),
		pc:       Word(0x200),
		mem:      &Memory{},
		video:    &Video{},
		keyboard: &Keyboard{},
	}
}

func (c *chip8) loadFonts() {
	mFontAddr := Word(0x000)
	for _, f := range fonts {
		c.mem.Write(mFontAddr, f)
		mFontAddr++
	}
}

func (c *chip8) Load(romFile string) error {

	f, err := os.Open(romFile)
	if err != nil {
		return err
	}
	defer f.Close()

	start := CPU_START
	for {

		b := make([]byte, 1)
		if _, err := f.Read(b); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		for _, b1 := range b {
			c.mem.Write(Word(start), b1)
			start++
		}
	}
}

func (c *chip8) Loop() {

	// set PROGRAM COUNTER to 0x200 (chip 8 memory start)
	c.pc = CPU_START
	// reset stack pointer
	c.sp = SP_START

	// load built-in font-set
	c.loadFonts()

	c.video.Init(512, 256)
	defer c.video.Close()

	for !rl.WindowShouldClose() && !c.stop {
		// execute single instruction
		c.execute_instruction()

		// emulate raylib event loop
		c.video.Draw([]byte{}, 0, 0)
	}
}

func (c *chip8) execute_instruction() {

	pc := c.pc
	var opcode Word
	hi := uint16(c.mem.Read(int(c.pc)))
	c.pc++
	lo := uint16(c.mem.Read(int(c.pc)))
	c.pc++
	opcode = Word((hi << 8) | (lo & 0xff))

	if opcode == 0 {
		c.pc -= 2
		return
	}

	fmt.Printf("PC=$%X OP=$%X\n", pc, opcode)

	switch opcode.High() & 0xf0 {

	case 0x00:
		switch opcode.Low() {
		default:
			c.pc -= 2
		case 0xE0:
			// 00E0 - CLS - Clear video screen
			c.cls(opcode)
		case 0xFD:
			// SCHIP - 00FD - EXIT - Exit interpreter
			c.stop = true
		case 0xEE:
			// 00EE - RTS - Return from Subroutine
			c.rts()
		case 0xFB:
			// SCHIP - 00FB - SCRR - Scroll the display right by 4 pixels
			c.scrr(opcode)
		case 0xFC:
			// SCHIP - 00FC - SCRL - Scroll the display left by 4 pixels
			c.scrl(opcode)
		case 0xFE:
			// SCHIP - 00FE - LORES - Disable high resolution
			c.video.extended = false
		case 0xFF:
			// SCHIP - 00FF - HIRES - Enable high resolution
			c.video.extended = true
		}
	case 0x10:
		// 1nnn - JUMP - Jump to Address
		c.jump(opcode)
	case 0x20:
		// 2nnn - CALL - Call Subroutine
		c.call(opcode)
	case 0x30:
		// 3snn - SKE - Skip if Register Equal Value
		c.ske(opcode)
	case 0x40:
		// 4snn - SKNE - Skip if Register Not Equal Value
		c.skne(opcode)
	case 0x50:
		// 5st0 - SKRE - Skip if Register Equal Register
		c.skre(opcode)
	case 0x60:
		// 6snn - LOAD - Load Register with Value
		c.load(opcode)
	case 0x70:
		// 7snn - ADD - Add value to register
		c.add(opcode)
	case 0x80:
		switch opcode.Low() & 0xF {
		case 0x0:
			// 8ts0 - MOVE - Move value between registers
			c.move(opcode)
		case 0x1:
			// 8ts1 - OR - Logical OR
			c.or(opcode)
		case 0x2:
			// 8ts2 - AND - Logical AND
			c.and(opcode)
		case 0x3:
			// 8ts3 - XOR - Logical XOR
			c.xor(opcode)
		case 0x4:
			// 8ts4 - ADDR - Add Register to Register with Overflow
			c.addr(opcode)
		case 0x5:
			// 8ts5 - SUB - Subtract Value from Register
			c.sub(opcode)
		case 0x6:
			// 8s06 - SHR - Shift Right
			c.shr(opcode)
		case 0x7:
			// 8s07 - SUBN - Subtract Register from Register
			c.subn(opcode)
		case 0xE:
			// 8s0E - SHL - Shift Left
			c.shl(opcode)
		default:
			break
		}
	case 0x90:
		// 9st0 - SKRNE - Skip if Register not Equal Register
		c.skrne(opcode)
	case 0xA0:
		// Annn - LOADI - Load Index
		c.load_i(opcode)
	case 0xB0:
		// Bnnn - JUMPI - Jump to Address with Index
		c.jump_i(opcode)
	case 0xC0:
		// Ctnn - RAND - Generate Random Number
		c.rnd(opcode)
	case 0xD0:
		// Dstn - DRAW - Draw Sprite
		c.draw(opcode)
	case 0xE0:
		switch opcode.Low() {
		case 0x9E:
			// Ex9E - SKP Vx - Skip next instruction if key with the value of Vx is pressed
			c.skp(opcode)
		case 0xA1:
			// ExA1 - SKNP Vx - Skip next instruction if key with the value of Vx is not pressed
			c.sknp(opcode)
		default:
			break
		}
	case 0xF0:
		switch opcode.Low() {
		case 0x07:
			// Ft07 - MOVED - Move Delay Register into Register
			c.move_d(opcode)
		case 0x0A:
			// Fx0A - LD Vx, K - Wait for a key press, store the value of the key in Vx
			c.keyd(opcode)
		case 0x15:
			// Fs15 - LOADD - Load Register into Delay Register
			c.load_d(opcode)
		case 0x18:
			// Fs18 - LOADS - Load Register into Sound Register
			c.loads(opcode)
		case 0x1E:
			// Fs1E - ADDI - Add Register into Index
			c.add_i(opcode)
		case 0x29:
			// Fs29 - LDSPR - Load Index with Sprite
			c.ldspr(opcode)
		case 0x33:
			// Fs33 - BCD - Store Binary Coded Decimal
			c.bcd(opcode)
		case 0x55:
			// Fs55 - STOR - Store Registers in Index
			c.stor(opcode)
		case 0x65:
			// Fs65 - READ - Read Stored Registers
			c.read(opcode)
		case 0x75:
			// Fn75 - SRPL - Stores the values from n number of registers into RPL
			c.srpl(opcode)
		case 0x85:
			// Fn85 - RRPL - Reads the values from n number of registers from RPL
			c.rrpl(opcode)
		default:
			break
		}
	}

	// decrement Delay Timer
	if c.dt > 0 {
		c.dt--
	}

	// decrement sound timer
	if c.st > 0 {
		c.st--
	}
}
