package main

import (
	"flag";
	"io";
	"os";
	"fmt";
)

//	Modified from https://openi.com/article/18/6/copying-files-go
func main() {
	a := false;
	b := flag.String("outputpathdata", a, "Where you want it to copy the binary.");
	r := flag.String("filenamedata", a, "What you want it to be named.");
	d := flag.String("miscdata", a, "What the original file is named.");
	e := flag.String("symboldata", "/", "Symbol data payload, usually just a forward slash.");
	f(d, b + e + r);
}
func f(p, q string) (int64, error) {
	g, h := os.Stat(p)
	if h != nil {
		return 0, h
	}
	if !g.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", p)
	}
	i, j := os.Open(p)
	if j != nil {
		return 0, j
	}
	defer i.Close()
	k, l := os.Create(q)
	if l != nil {
		return 0, l
	}
	defer k.Close()
	n, o := io.Copy(k, i)
	return n, o
}
