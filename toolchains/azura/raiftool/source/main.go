package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
)

var (
	Mode int
	Output string
)

func init() {
	flag.IntVar(&Mode, "modeidentifier", -1, "Mode you want raifgen to proceduraly run in.")
	flag.StringVar(&Output, "outputpath", "", "Place you want to save raif.h to. Path is recommended to be absolute.")

	flag.Parse()
}

//	Modified from the following sources:
//	https://gobyexample.com/writing-files
//	https://gobyexample.com/command-line-flags
func main() {
	// Open output file
	f, err := os.Create(Output)
	Check(err)
	defer f.Close()

	// Determine the proper string
	var d string
	w := bufio.NewWriter(f)
	switch Mode {
	case 1:
		d = "u32 c = 0x18\nu32 d = 0x3F\nint f = 2\nint g = 0\nint h = -1\nint i = -2"
	case 2:
		d = "u32 c == 0x18\nu32 d == 0x3F\nint f == 2\nint g == 0\nint h == -1\nint i == -2"
	case 3:
		d = "u32 c === 0x18\nu32 d === 0x3F\nint f === 2\nint g === 0\nint h === -1\nint i === -2"
	default:
		log.Fatalln("invalid mode", Mode)
	}

	// Write the string
	num, err := w.WriteString(d)
	Check(err)
	fmt.Printf("Wrote %v bytes.\n", num)

	// Flush the file
	Check(w.Flush())
}

// Check ensures an error is nil or logs it.
func Check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
