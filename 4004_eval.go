package main

import (
  // "fmt"
  "bufio"
)

// NOTE: OK
func (x *CPU) nop() {
}

func (x *CPU) jcn(cond byte, addr2 byte, addr1 byte) {
}

func (x *CPU) fim(reg byte, data2 byte, data1 byte) {
}

func (x *CPU) fin(reg byte) {
}

func (x *CPU) jin(reg byte) {
}

// NOTE: OK
func (x *CPU) jun(addr3 byte, addr2 byte, addr1 byte) {
  addr := uint16(addr3) << 8 | uint16(addr2) << 4 | uint16(addr1) << 0
  x.rom_addr = addr
}

func (x *CPU) jms(addr3 byte, addr2 byte, addr1 byte) {
}

// NOTE: OK
func (x *CPU) inc(reg byte) {
  x.regs[reg]++
}

func (x *CPU) isz(reg byte, addr2 byte, addr1 byte) {
}

func (x *CPU) add(reg byte) {
}

func (x *CPU) sub(reg byte) {
}

// NOTE: OK
func (x *CPU) ld(reg byte) {
  x.accum = x.regs[reg]
}

// NOTE: OK
func (x *CPU) xch(reg byte) {
  x.accum, x.regs[reg] = x.regs[reg], x.accum
}

func (x *CPU) bbl(data byte) {
}

// NOTE: OK
func (x *CPU) ldm(data byte) {
  x.accum = data
}

// NOTE: OK
func (x *CPU) clb() {
  x.accum = 0
  x.carry = 0
}

// NOTE: OK
func (x *CPU) clc() {
  x.carry = 0
}

// NOTE: OK
func (x *CPU) iac() {
  x.accum++
}

// NOTE: OK
func (x *CPU) cmc() {
  x.carry = ^x.carry
}

// NOTE: OK
func (x *CPU) cma() {
  x.accum = ^x.accum
}

// NOTE: OK
func (x *CPU) ral() {
  x.carry <<= 1
  x.accum <<= 1
}

// NOTE: OK
func (x *CPU) rar() {
  x.carry >>= 1
  x.accum >>= 1
}

func (x *CPU) tcc() {
}

// NOTE: OK
func (x *CPU) dac() {
  x.accum--
}

func (x *CPU) tcs() {
}

// NOTE: OK
func (x *CPU) stc() {
  x.carry = 1
}

func (x *CPU) daa() {
}

func (x *CPU) kbp() {
}

func (x *CPU) dcl() {
}

func (x *CPU) src(reg byte) {
}

func (x *CPU) wrm() {
}

func (x *CPU) wmp() {
}

func (x *CPU) wrr() {
}

func (x *CPU) wpm() {
}

func (x *CPU) wr0() {
}

func (x *CPU) wr1() {
}

func (x *CPU) wr2() {
}

func (x *CPU) wr3() {
}

func (x *CPU) sbm() {
}

func (x *CPU) rdm() {
}

func (x *CPU) rdr() {
}

func (x *CPU) adm() {
}

func (x *CPU) rd0() {
}

func (x *CPU) rd1() {
}

func (x *CPU) rd2() {
}

func (x *CPU) rd3() {
}

func (x *CPU) evalMachine(op byte, r *bufio.Reader) {
  opr := (op & 0xF0) >> 4
  opa := (op & 0x0F) >> 0
  switch opr {
  case 0x0:
    if opa != 0x0 {
      panic("invalid operation")
    }
    x.nop()
  case 0x1:
    addr, _ := r.ReadByte()
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    x.jcn(opa, addrr, addra)
  case 0x2:
    reg := (opa & 0xE) >> 1
    flag := (opa & 0x1) >> 0
    if flag == 0 {
      data, _ := r.ReadByte()
      datar := (data & 0xF0) >> 4
      dataa := (data & 0x0F) >> 0
      x.fim(reg, datar, dataa)
    } else {
      x.src(reg)
    }
  case 0x3:
    reg := (opa & 0xE) >> 1
    flag := (opa & 0x1) >> 0
    if flag == 0 {
      x.fin(reg)
    } else {
      x.jin(reg)
    }
  case 0x4:
    addr, _ := r.ReadByte()
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    x.jun(opa, addrr, addra)
  case 0x5:
    addr, _ := r.ReadByte()
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    x.jms(opa, addrr, addra)
  case 0x6: x.inc(opa)
  case 0x7:
    addr, _ := r.ReadByte()
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    x.isz(opa, addrr, addra)
  case 0x8: x.add(opa)
  case 0x9: x.sub(opa)
  case 0xA: x.ld(opa)
  case 0xB: x.xch(opa)
  case 0xC: x.bbl(opa)
  case 0xD: x.ldm(opa)
  }
}

func (x *CPU) evalIORAM(op byte) {
  switch op {
  case 0xE0: x.wrm()
  case 0xE1: x.wmp()
  case 0xE2: x.wrr()
  case 0xE3: x.wpm()
  case 0xE4: x.wr0()
  case 0xE5: x.wr1()
  case 0xE6: x.wr2()
  case 0xE7: x.wr3()
  case 0xE8: x.sbm()
  case 0xE9: x.rdm()
  case 0xEA: x.rdr()
  case 0xEB: x.adm()
  case 0xEC: x.rd0()
  case 0xED: x.rd1()
  case 0xEE: x.rd2()
  case 0xEF: x.rd3()
  }
}

func (x *CPU) evalAccum(op byte) {
  switch op {
  case 0xF0: x.clb()
  case 0xF1: x.clc()
  case 0xF2: x.iac()
  case 0xF3: x.cmc()
  case 0xF4: x.cma()
  case 0xF5: x.ral()
  case 0xF6: x.rar()
  case 0xF7: x.tcc()
  case 0xF8: x.dac()
  case 0xF9: x.tcs()
  case 0xFA: x.stc()
  case 0xFB: x.daa()
  case 0xFC: x.kbp()
  case 0xFD: x.dcl()
  }
}

func (x *CPU) eval(r *bufio.Reader) {
  for op, err := r.ReadByte(); err == nil; op, err = r.ReadByte() {
    opr := (op & 0xF0) >> 4
    switch opr {
    case 0xE:
      x.evalIORAM(op)
    case 0xF:
      x.evalAccum(op)
    default:
      x.evalMachine(op, r)
    }
  }
}
