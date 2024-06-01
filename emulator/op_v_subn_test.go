package emulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_chip8_subn(t *testing.T) {
	type args struct {
		opcode Word
	}
	tests := []struct {
		name string
		c    func() *chip8
		args args
		want func(t *testing.T, c *chip8)
	}{
		{
			name: "subn borrow",
			c: func() *chip8 {
				reg := make([]uint8, 16)
				reg[0] = 0x5
				reg[1] = 0x4
				return &chip8{
					reg: reg,
				}
			},
			args: args{
				opcode: 0x8017,
			},
			want: func(t *testing.T, c *chip8) {
				assert.Equal(t, uint8(0xFF), c.reg[0x0])
				assert.Equal(t, uint8(0), c.reg[0xF])
			},
		},
		{
			name: "subn not borrow",
			c: func() *chip8 {
				reg := make([]uint8, 16)
				reg[0] = 0x6
				reg[1] = 0x7
				return &chip8{
					reg: reg,
				}
			},
			args: args{
				opcode: 0x8017,
			},
			want: func(t *testing.T, c *chip8) {
				assert.Equal(t, uint8(0x1), c.reg[0x0])
				assert.Equal(t, uint8(1), c.reg[0xF])
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.c()
			c.subn(tt.args.opcode)
			tt.want(t, c)
		})
	}
}
