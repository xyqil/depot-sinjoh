package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	Mode   string
	Output string
)

func init() {
	flag.StringVar(&Mode, "modeidentifier", "", "Mode you want boolgen to proceduraly run in.")
	flag.StringVar(&Output, "outputpath", "", "Place you want to save zuul.h to. Path is recommended to be absolute.")

	flag.Parse()
}

//	Modified from the following sources:
//	https://gobyexample.com/writing-files
//	https://gobyexample.com/command-line-flags
func main() {
	// Open file for output.
	f, err := os.Open(Output)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)

	// Determine proper string.
	// Configurable to meet all three equality types for strings!
	var d string
	switch c {
	case 1:
		d = "bool f = true\nbool g = false"
	case 2:
		d = "bool f == true\nbool g == false"
	case 3:
		d = "bool f === true\nbool g === false"
	default:
		log.Fatalln("invalid c", c)
	}
	// TODO: Maybe rename the one-letters, if it doesn't save RAM/binary space?
	// Write the string
	if _, err = w.WriteString(d); err != nil {
		log.Fatalln("error writing", err)
	}

	// Flush the file
	if err = w.Flush(); err != nil {
		log.Fatalln("error flushing", err)
	}
}