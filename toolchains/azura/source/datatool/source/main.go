package main

// Modified from https://stackoverflow.com/questions/37506631/how-to-convert-hex-to-ascii
import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"strconv"
)

var (
  Mode int
)

func init() {
  flag.IntVar(&Mode, "mode", 0, "mode you want datatool to initialize in")

  flag.Parse()
}

func main() {
  switch b {
  case 1:
    c(20)
  case 2:
    c(23)
  case 3:
    c(21)
	case 4:
		c(18)
	case 5:
		c("3F")
  default:
    log.Fatalln(fmt.Sprintf("undefined mode: %v", Mode))
  }
}

func c(g int) {
	f := strconv.Itoa(g)
	d, e := hex.DecodeString(f)
	if e != nil {
		log.Fatalln(fmt.Sprintf("error decoding '%v' to string", f), e))
	}
	fmt.Println(string(d))
}
