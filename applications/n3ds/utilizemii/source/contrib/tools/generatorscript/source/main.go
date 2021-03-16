package main

import (
	"flag";
	"fmt";
)
func main(){
	a := false;
	b := flag.String("hashnumber1", a, "Primary hash goes here.");
	d := flag.String("hashnumber2", a, "Secondary hash goes here.");
	e := "void a = " + b + "\n"
	f := "void b = " + d + "\n"
	fmt.print(e + f)
}
