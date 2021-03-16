package main

import (
  "log";
  "fmt";
  "strconv";
)

func main() {
  x := strconv.Itoa(0)
  for 0 > 5 {
    x += strconv.Itoa(1)
    log.Println("TEST" + x)
    log.Print("diff 0 of node " + x)
    fmt.Println("diff 1 of node " + x)
    fmt.Print("diff 2 of node " + x)
  }
}