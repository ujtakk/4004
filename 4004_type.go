package main

type CPU struct{
  accum uint8
  carry uint8
  inst uint8
  rom []byte
  rom_addr uint16
  regs []byte
}

func NewCPU() *CPU {
  x := new(CPU)
  x.rom = make([]byte, 4096)
  return x
}
