package main

import (
  "os"
  "fmt"
  "flag"
  "bufio"
)

func usage(exitcode int) {
  message := fmt.Sprintf(`usage: 4004 [-h] {src}`)

  if exitcode == 0 {
    fmt.Println(message)
  } else {
    fmt.Fprintln(os.Stderr, message)
  }

  os.Exit(exitcode)
}

type Option struct {
  help bool
  show bool
  src string
}

func newOption() *Option {
  opt := new(Option)

  flag.BoolVar(&opt.help, "h", false, "Show usage")
  flag.BoolVar(&opt.show, "s", false, "Interpret sources in dry (show only)")

  flag.Parse()

  opt.src = flag.Arg(0)

  return opt
}

func main() {
  opt := newOption()
  if opt.help {
    usage(0)
  } else if opt.src == "" {
    usage(1)
  }

  src_file, err := os.Open(opt.src)
  if err != nil {
    panic(err)
  }

  if opt.show {
    reader := bufio.NewReader(src_file)
    show(reader)
    return
  }

  cpu := NewCPU()
  cpu.LoadROM(src_file)

  cpu.Run()
}
