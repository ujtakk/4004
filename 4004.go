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
  src string
}

func newOption() *Option {
  opt := new(Option)

  flag.BoolVar(&opt.help, "h", false, "Show usage")

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

  file, err := os.Open(opt.src)
  if err != nil {
    panic(err)
  }

  reader := bufio.NewReader(file)
  show(reader)
}
