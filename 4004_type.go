package main

const (
  ROM_SIZE = 4096
)

type CPU struct{
  accum uint8
  carry uint8
  inst uint8
  rom []byte
  regs []byte
  pcounter uint16
  spointer uint8
  test uint8
}

func NewCPU() *CPU {
  x := new(CPU)
  x.rom = make([]byte, ROM_SIZE)
  return x
}

func (x *CPU) LoadROM() {
}

func (x *CPU) SaveRAM(path string) {
}

func (x *CPU) Exec() {
}
