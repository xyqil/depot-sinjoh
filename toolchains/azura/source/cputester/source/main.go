package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/millefalcon/go-subprocess/subprocess"
)

var (
	// TODO: What are these
	// Oh do a base64 to ascii conversion, combine both side by side too also :)
	a = "Y29uc29sZS5sb2coIkF1dGhvcml6ZWQgdG8g"
	b = "dXNlIDYxMDBtXCdccyBTdHJlYW0gQm90Iik="
)

//	Modified from the following sources:
//	https://stackoverflow.com/questions/37506631/how-to-convert-hex-to-ascii
//	https://www.base64decoder.io/golang/
//	https://gobyexample.com/writing-files
func main() {
	f, err := os.Create("escapetest.js")
	if err != nil {
		log.Fatalln("error creating escapetest.js", err)
	}

	w := bufio.NewWriter(f)
	
	num, err := w.WriteString(DecodeString(a + b))
	if err != nil {
		log.Fatalln("error writing string", err)
	}

	// this doesn't seem necessary
	// Describe(i, k)
	//Eh, saves if err = nil statements tbh, or at least makes main not really cluttered a bit.
	// I guess, but i think it's better to be more explicit by checking errors sooner
	//Yeah, I agree.
	log.Printf("Wrote %v bytes\n", num)
	
	// Flush the file
	if err = w.Flush(); err != nil {
		log.Fatalln("error flushing file", err)
	}

	// Run node
	f := subprocess.Popen("node escapetest.js") //escapetest.js is commonly refered to as test.js by 6100m, the name of the techdemo if you will.
	log.Println(f.Stdout.Read())
}

/*
// Describe pretty-prints a helpful error message.
func Describe(a, b error) {
	if a != nil {
		log.Fatalln("error, a variable was defined as: %v", a.Error())
	}
	if b != nil {
		log.Fatalln("error, a variable was defined as: %v", b.Error())
	}
}*/

// DecodeString decodes base64-encoded strings with the t variable.
func DecodeString(t string) string {
	v, err := base64.StdEncoding.DecodeString(t)
	if err != nil {
		log.Fatalln("error decoding string", err)
	}

	return string(v)
}

