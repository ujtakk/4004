package main

type CPU struct{
  accum uint8
  inst uint8
  rom []byte
}

func NewCPU() *CPU {
  x := new(CPU)
  return x
}
