package main

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

func (x *CPU) LoadROM() {
}

func (x *CPU) SaveRAM(path string) {
}

func (x *CPU) Exec() {
}
