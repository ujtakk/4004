package main

import (
  "fmt"
  "bufio"
)

func showMachine(op byte, r *bufio.Reader) {
  opr := (op & 0xF0) >> 4
  opa := (op & 0x0F) >> 0
  switch opr {
  case 0x0:
    if opa != 0x0 {
      panic("invalid operation")
    }
    fmt.Printf("NOP\n")
  case 0x1:
    addr, _ := r.ReadByte()
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    fmt.Printf("JCN %01x %01x %01x\n", opa, addrr, addra)
  case 0x2:
    reg := (opa & 0xE) >> 1
    flag := (opa & 0x1) >> 0
    if flag == 0 {
      data, _ := r.ReadByte()
      datar := (data & 0xF0) >> 4
      dataa := (data & 0x0F) >> 0
      fmt.Printf("FIM %01x %01x %01x\n", reg, datar, dataa)
    } else {
      fmt.Printf("SRC %01x\n", reg)
    }
  case 0x3:
    reg := (opa & 0xE) >> 1
    flag := (opa & 0x1) >> 0
    if flag == 0 {
      fmt.Printf("FIN %01x\n", reg)
    } else {
      fmt.Printf("JIN %01x\n", reg)
    }
  case 0x4:
    addr, _ := r.ReadByte()
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    fmt.Printf("JUN %01x %01x %01x\n", opa, addrr, addra)
  case 0x5:
    addr, _ := r.ReadByte()
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    fmt.Printf("JMS %01x %01x %01x\n", opa, addrr, addra)
  case 0x6:
    fmt.Printf("INC %01x\n", opa)
  case 0x7:
    addr, _ := r.ReadByte()
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    fmt.Printf("ISZ %01x %01x %01x\n", opa, addrr, addra)
  case 0x8:
    fmt.Printf("ADD %01x\n", opa)
  case 0x9:
    fmt.Printf("SUB %01x\n", opa)
  case 0xA:
    fmt.Printf("LD  %01x\n", opa)
  case 0xB:
    fmt.Printf("XCH %01x\n", opa)
  case 0xC:
    fmt.Printf("BBL %01x\n", opa)
  case 0xD:
    fmt.Printf("LDM %01x\n", opa)
  }
}

func showIORAM(op byte) {
  // opr := (op & 0xF0) >> 4
  opa := (op & 0x0F) >> 0
  switch opa {
  case 0x0: fmt.Println("WRM")
  case 0x1: fmt.Println("WMP")
  case 0x2: fmt.Println("WRR")
  case 0x3: fmt.Println("WPM")
  case 0x4: fmt.Println("WR0")
  case 0x5: fmt.Println("WR1")
  case 0x6: fmt.Println("WR2")
  case 0x7: fmt.Println("WR3")
  case 0x8: fmt.Println("SBM")
  case 0x9: fmt.Println("RDM")
  case 0xA: fmt.Println("RDR")
  case 0xB: fmt.Println("ADM")
  case 0xC: fmt.Println("RD0")
  case 0xD: fmt.Println("RD1")
  case 0xE: fmt.Println("RD2")
  case 0xF: fmt.Println("RD3")
  }
}

func showAccum(op byte) {
  // opr := (op & 0xF0) >> 4
  opa := (op & 0x0F) >> 0
  switch opa {
  case 0x0: fmt.Println("CLB")
  case 0x1: fmt.Println("CLC")
  case 0x2: fmt.Println("IAC")
  case 0x3: fmt.Println("CMC")
  case 0x4: fmt.Println("CMA")
  case 0x5: fmt.Println("RAL")
  case 0x6: fmt.Println("RAR")
  case 0x7: fmt.Println("TCC")
  case 0x8: fmt.Println("DAC")
  case 0x9: fmt.Println("TCS")
  case 0xA: fmt.Println("STC")
  case 0xB: fmt.Println("DAA")
  case 0xC: fmt.Println("KBP")
  case 0xD: fmt.Println("DCL")
  }
}

func show(r *bufio.Reader) {
  for op, err := r.ReadByte(); err == nil; op, err = r.ReadByte() {
    opr := (op & 0xF0) >> 4
    switch opr {
    case 0xE:
      showIORAM(op)
    case 0xF:
      showAccum(op)
    default:
      showMachine(op, r)
    }
  }
}
