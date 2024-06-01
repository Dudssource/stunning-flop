package emulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_chip8_rnd(t *testing.T) {
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
			name: "rnd",
			c: func() *chip8 {
				reg := make([]uint8, 16)
				reg[0xA] = 0x0
				return &chip8{
					reg: reg,
				}
			},
			args: args{
				opcode: 0xCAFF,
			},
			want: func(t *testing.T, c *chip8) {
				assert.Greater(t, uint8(0), c.reg[0xA])
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.c()
			c.rnd(tt.args.opcode)
			tt.want(t, c)
		})
	}
}
