package main

import (
	"flag"
	"os"

	"github.com/Dudssource/stunning-flop/emulator"
)

func writeROM(file string) {

	buf := []byte{
		// LOAD V0, A
		0x60, 0x0A,
		// LOAD V1, 5
		0x61, 0x05,
		// LOAD V3, 9
		0x63, 0x09,
		// LOAD I V3
		0xF3, 0x29,
		// DRAW X=V0, Y=V1, N=5
		0xD0, 0x15,
		// LOAD V0, 14
		0x60, 0x14,
		// LOAD V1, 5
		0x61, 0x05,
		// LOAD V3, 8
		0x63, 0x08,
		// LOAD I V3
		0xF3, 0x29,
		// DRAW X=V0, Y=V1, N=5
		0xD0, 0x15,
		// FA0A
		0xFA, 0x0A,
	}

	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(buf)

	if err != nil {
		panic(err)
	}
	f.Sync()
}

func main() {

	flag.Args()

	file := flag.String("f", "", "ROM `file` location")
	assembler := flag.String("m", "", "Assembler `mode`")
	flag.Parse()

	if file == nil || *file == "" {
		flag.Usage()
		os.Exit(1)
		return
	}

	if assembler != nil && *assembler == "asm" {
		writeROM(*file)
		return
	}

	cpu := emulator.Cpu()
	if err := cpu.Load(*file); err != nil {
		panic(err)
	}

	cpu.Loop()
}
