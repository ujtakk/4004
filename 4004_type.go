package main

import (
	"io"
	"io/ioutil"
)

const (
	ROM_SIZE   = 4096
	RAM_LINE   = 4
	RAM_SIZE   = 20
	NUM_REGS   = 16
	STACK_SIZE = 4
)

type CPU struct {
	io       uint8
	out      uint8
	test     uint8
	ctrl     uint8
	ram_ctrl uint8
	accum    uint8
	carry    uint8
	spointer uint8
	pcounter *uint16
	regs     []byte
	rom      []byte
	rams     [][]byte
	stack    []uint16
}

func NewCPU() *CPU {
	x := new(CPU)

	x.regs = make([]byte, NUM_REGS)

	x.rom = make([]byte, ROM_SIZE)

	x.rams = make([][]byte, RAM_LINE)
	for i := 0; i < RAM_LINE; i++ {
		x.rams[i] = make([]byte, RAM_SIZE)
	}

	x.stack = make([]uint16, STACK_SIZE)
	x.pcounter = &x.stack[0]

	return x
}

func (x *CPU) LoadROM(r io.Reader) {
	data, _ := ioutil.ReadAll(r)
	if len(data) > ROM_SIZE {
		panic("exceeded the ROM size")
	}

	copy(x.rom, data)
}

func (x *CPU) Run() {
	for *x.pcounter < ROM_SIZE {
		inst := x.rom[*x.pcounter]
		x.eval(inst)
		*x.pcounter++
	}
}

func (x *CPU) SaveRAM(w io.Writer) {
	for i := 0; i < RAM_LINE; i++ {
		n, err := w.Write(x.rams[i])
		if err != nil {
			panic(err)
		} else if n != RAM_SIZE {
			panic("data was not saved correctly")
		}
	}
}
