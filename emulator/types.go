package emulator

type Word uint16

func (w Word) High() uint8 {
	return uint8(w & 0xff00 >> 8)
}

func (w Word) Low() uint8 {
	return uint8(w & 0x00ff)
}

func (w Word) Word() uint16 {
	return uint16(w)
}

type Memory struct {
	mem [4096]uint8
}

func (m *Memory) Read(address int) uint8 {
	return m.mem[address]
}

func (m *Memory) Write(address Word, value uint8) {
	m.mem[address] = value
}
