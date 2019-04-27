package main

import (
  "io"
  "io/ioutil"
)

const (
  ROM_SIZE = 4096
)

type CPU struct{
  test uint8
  ctrl uint8
  ram_ctrl uint8
  accum uint8
  carry uint8
  spointer uint8
  pcounter *uint16
  rom []byte
  regs []byte
  rams [][]byte
  stack []uint16
}

func NewCPU() *CPU {
  x := new(CPU)
  x.pcounter = &x.stack[0]
  x.rom = make([]byte, ROM_SIZE)
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
}
