package main

// NOTE: OK
func (x *CPU) nop() {
}

// NOTE: OK
func (x *CPU) jcn(cond byte, addr2 byte, addr1 byte) {
  switch {
  case (cond & 001 == 001) && (x.test == 0): fallthrough
  case (cond & 002 == 002) && (x.carry == 1): fallthrough
  case (cond & 004 == 004) && (x.accum == 0): fallthrough
  case (cond & 011 == 011) && (x.test != 0): fallthrough
  case (cond & 012 == 012) && (x.carry != 1): fallthrough
  case (cond & 014 == 014) && (x.accum != 0):
    addr := uint16(addr2) << 4 | uint16(addr1) << 0
    *x.pcounter = (*x.pcounter & 0x0F00) | addr
  }
}

// NOTE: OK
func (x *CPU) fim(reg byte, data2 byte, data1 byte) {
  reg_addr := reg << 1
  x.regs[reg_addr+0] = data2
  x.regs[reg_addr+1] = data1
}

// NOTE: OK
func (x *CPU) fin(reg byte) {
  reg_addr := reg << 1
  rom_addr := uint16(x.regs[0]) << 4 | uint16(x.regs[1]) << 0
  data := x.rom[rom_addr]
  x.regs[reg_addr+0] = (data & 0xF0) >> 4
  x.regs[reg_addr+1] = (data & 0x0F) >> 0
}

// NOTE: OK
func (x *CPU) jin(reg byte) {
  reg_addr := reg << 1
  addr2 := x.regs[reg_addr+0]
  addr1 := x.regs[reg_addr+1]
  addr := uint16(addr2) << 4 | uint16(addr1) << 0
  *x.pcounter = (*x.pcounter & 0x0F00) | addr
}

// NOTE: OK
func (x *CPU) jun(addr3 byte, addr2 byte, addr1 byte) {
  addr := uint16(addr3) << 8 | uint16(addr2) << 4 | uint16(addr1) << 0
  *x.pcounter = addr
}

// NOTE: OK
func (x *CPU) jms(addr3 byte, addr2 byte, addr1 byte) {
  x.spointer++
  x.pcounter = &x.stack[x.spointer]
}

// NOTE: OK
func (x *CPU) inc(reg byte) {
  x.regs[reg] = (x.regs[reg] + 1) & 0x0F
}

// NOTE: OK
func (x *CPU) isz(reg byte, addr2 byte, addr1 byte) {
  x.regs[reg] = (x.regs[reg] + 1) & 0x0F
  if x.regs[reg] != 0 {
    addr := uint16(addr2) << 4 | uint16(addr1) << 0
    *x.pcounter = addr
  }
}

// NOTE: OK
func (x *CPU) add(reg byte) {
  result := x.accum + x.carry + x.regs[reg]
  x.accum = (result & 0x0F) >> 0
  x.carry = (result & 0xF0) >> 4
}

// NOTE: OK
func (x *CPU) sub(reg byte) {
  // result := x.accum - x.regs[reg]
  result := x.accum + x.carry + (x.regs[reg] ^ 0x0F)
  x.accum = (result & 0x0F) >> 0
  x.carry = (result & 0xF0) >> 4
}

// NOTE: OK
func (x *CPU) ld(reg byte) {
  x.accum = x.regs[reg]
}

// NOTE: OK
func (x *CPU) xch(reg byte) {
  x.accum, x.regs[reg] = x.regs[reg], x.accum
}

// NOTE: OK
func (x *CPU) bbl(data byte) {
  x.spointer--
  x.pcounter = &x.stack[x.spointer]
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
  x.carry = x.carry ^ 0x0F
}

// NOTE: OK
func (x *CPU) cma() {
  x.accum = x.accum ^ 0x0F
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

// NOTE: OK
func (x *CPU) tcc() {
  switch x.carry {
  case 0x1: x.accum = 0x1
  case 0x0: x.accum = 0x0
  }
  x.carry = 0
}

// NOTE: OK
func (x *CPU) dac() {
  x.accum--
}

// NOTE: OK
func (x *CPU) tcs() {
  switch x.carry {
  case 0x1: x.accum = 0x9
  case 0x0: x.accum = 0xA
  }
  x.carry = 0
}

// NOTE: OK
func (x *CPU) stc() {
  x.carry = 1
}

// NOTE: OK
func (x *CPU) daa() {
  if 0xA <= x.accum && x.accum <= 0xF {
    x.accum += 6
  }
}

// NOTE: OK
func (x *CPU) kbp() {
  switch x.accum {
  case 0x0: x.accum = 0x0
  case 0x1: x.accum = 0x1
  case 0x2: x.accum = 0x2
  case 0x4: x.accum = 0x3
  case 0x8: x.accum = 0x4
  default: x.accum = 0xF
  }
}

// NOTE: OK
func (x *CPU) dcl() {
  x.ctrl = x.accum
}

// NOTE: OK
func (x *CPU) src(reg byte) {
  reg_addr := reg << 1
  addr2 := x.regs[reg_addr+0]
  addr1 := x.regs[reg_addr+1]
  addr := uint8(addr2) << 4 | uint8(addr1) << 0
  x.ram_ctrl = addr
}

// NOTE: OK
func (x *CPU) wrm() {
  x.rams[x.ctrl][x.ram_ctrl] = x.accum
}

// NOTE: OK
func (x *CPU) wmp() {
  x.out = x.accum
}

// NOTE: OK
func (x *CPU) wrr() {
  x.io = x.accum
}

// NOTE: OK
func (x *CPU) wpm() {
  panic("not implemented for 4004")
}

// NOTE: OK
func (x *CPU) wr0() {
  x.rams[x.ctrl][16] = x.accum
}

// NOTE: OK
func (x *CPU) wr1() {
  x.rams[x.ctrl][17] = x.accum
}

// NOTE: OK
func (x *CPU) wr2() {
  x.rams[x.ctrl][18] = x.accum
}

// NOTE: OK
func (x *CPU) wr3() {
  x.rams[x.ctrl][19] = x.accum
}

// NOTE: OK
func (x *CPU) sbm() {
  // result := x.accum - x.rams[x.ctrl][x.ram_ctrl]
  result := x.accum + x.carry + (x.rams[x.ctrl][x.ram_ctrl] ^ 0x0F)
  x.accum = (result & 0x0F) >> 0
  x.carry = (result & 0xF0) >> 4
}

// NOTE: OK
func (x *CPU) rdm() {
  x.accum = x.rams[x.ctrl][x.ram_ctrl]
}

// NOTE: OK
func (x *CPU) rdr() {
  x.accum = x.io
}

// NOTE: OK
func (x *CPU) adm() {
  result := x.accum + x.carry + x.rams[x.ctrl][x.ram_ctrl]
  x.accum = (result & 0x0F) >> 0
  x.carry = (result & 0xF0) >> 4
}

// NOTE: OK
func (x *CPU) rd0() {
  x.accum = x.rams[x.ctrl][16]
}

// NOTE: OK
func (x *CPU) rd1() {
  x.accum = x.rams[x.ctrl][17]
}

// NOTE: OK
func (x *CPU) rd2() {
  x.accum = x.rams[x.ctrl][18]
}

// NOTE: OK
func (x *CPU) rd3() {
  x.accum = x.rams[x.ctrl][19]
}

func (x *CPU) evalMachine(inst byte) {
  opr := (inst & 0xF0) >> 4
  opa := (inst & 0x0F) >> 0
  switch opr {
  case 0x0:
    if opa != 0x0 {
      panic("invalid operation")
    }
    x.nop()
  case 0x1:
    *x.pcounter++
    addr := x.rom[*x.pcounter]
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    x.jcn(opa, addrr, addra)
  case 0x2:
    reg := (opa & 0xE) >> 1
    flag := (opa & 0x1) >> 0
    if flag == 0 {
      *x.pcounter++
      data := x.rom[*x.pcounter]
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
    *x.pcounter++
    addr := x.rom[*x.pcounter]
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    x.jun(opa, addrr, addra)
  case 0x5:
    *x.pcounter++
    addr := x.rom[*x.pcounter]
    addrr := (addr & 0xF0) >> 4
    addra := (addr & 0x0F) >> 0
    x.jms(opa, addrr, addra)
  case 0x6: x.inc(opa)
  case 0x7:
    *x.pcounter++
    addr := x.rom[*x.pcounter]
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

func (x *CPU) evalIORAM(inst byte) {
  switch inst {
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

func (x *CPU) evalAccum(inst byte) {
  switch inst {
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

func (x *CPU) eval(inst byte) {
  opr := (inst & 0xF0) >> 4
  switch opr {
  case 0xE:
    x.evalIORAM(inst)
  case 0xF:
    x.evalAccum(inst)
  default:
    x.evalMachine(inst)
  }
}
