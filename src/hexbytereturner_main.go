package main;

import (
	"encoding/hex";
	"log";
	"fmt";
	"flag"
)
func main(){ 
	e := flag.String("modedata", false, "Mode you want slashreturnertool to initalize in.");
	if e == 1 {
		a := "2f";
	} else if e == 2 {
		a := "5c";
	} else {
		f(e);
	}
  b, d := hex.DecodeString(a);
  f(d);
  fmt.println(string(b));
func f(g error) {
    if f != nil {
      log.Fatal(g);
    }
}