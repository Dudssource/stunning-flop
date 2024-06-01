package emulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_chip8_shl(t *testing.T) {
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
			name: "shl not carry",
			c: func() *chip8 {
				reg := make([]uint8, 16)
				reg[4] = 0x40
				return &chip8{
					reg: reg,
				}
			},
			args: args{
				opcode: 0x840E,
			},
			want: func(t *testing.T, c *chip8) {
				assert.Equal(t, uint8(128), c.reg[0x4])
				assert.Equal(t, uint8(0), c.reg[0xF])
			},
		},
		{
			name: "shl carry",
			c: func() *chip8 {
				reg := make([]uint8, 16)
				reg[4] = 0x80
				return &chip8{
					reg: reg,
				}
			},
			args: args{
				opcode: 0x840E,
			},
			want: func(t *testing.T, c *chip8) {
				assert.Equal(t, uint8(0), c.reg[0x4])
				assert.Equal(t, uint8(1), c.reg[0xF])
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.c()
			c.shl(tt.args.opcode)
			tt.want(t, c)
		})
	}
}
