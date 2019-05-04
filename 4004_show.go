package main

import (
	"bufio"
	"fmt"
)

func showMachine(inst byte, r *bufio.Reader) {
	opr := (inst & 0xF0) >> 4
	opa := (inst & 0x0F) >> 0
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

func showIORAM(inst byte) {
	switch inst {
	case 0xE0:
		fmt.Println("WRM")
	case 0xE1:
		fmt.Println("WMP")
	case 0xE2:
		fmt.Println("WRR")
	case 0xE3:
		fmt.Println("WPM")
	case 0xE4:
		fmt.Println("WR0")
	case 0xE5:
		fmt.Println("WR1")
	case 0xE6:
		fmt.Println("WR2")
	case 0xE7:
		fmt.Println("WR3")
	case 0xE8:
		fmt.Println("SBM")
	case 0xE9:
		fmt.Println("RDM")
	case 0xEA:
		fmt.Println("RDR")
	case 0xEB:
		fmt.Println("ADM")
	case 0xEC:
		fmt.Println("RD0")
	case 0xED:
		fmt.Println("RD1")
	case 0xEE:
		fmt.Println("RD2")
	case 0xEF:
		fmt.Println("RD3")
	}
}

func showAccum(inst byte) {
	switch inst {
	case 0xF0:
		fmt.Println("CLB")
	case 0xF1:
		fmt.Println("CLC")
	case 0xF2:
		fmt.Println("IAC")
	case 0xF3:
		fmt.Println("CMC")
	case 0xF4:
		fmt.Println("CMA")
	case 0xF5:
		fmt.Println("RAL")
	case 0xF6:
		fmt.Println("RAR")
	case 0xF7:
		fmt.Println("TCC")
	case 0xF8:
		fmt.Println("DAC")
	case 0xF9:
		fmt.Println("TCS")
	case 0xFA:
		fmt.Println("STC")
	case 0xFB:
		fmt.Println("DAA")
	case 0xFC:
		fmt.Println("KBP")
	case 0xFD:
		fmt.Println("DCL")
	}
}

func show(r *bufio.Reader) {
	for inst, err := r.ReadByte(); err == nil; inst, err = r.ReadByte() {
		opr := (inst & 0xF0) >> 4
		switch opr {
		case 0xE:
			showIORAM(inst)
		case 0xF:
			showAccum(inst)
		default:
			showMachine(inst, r)
		}
	}
}
