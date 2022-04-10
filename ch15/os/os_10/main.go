package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		s += " "
		if *NewLine {
			s += Newline
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)

}
