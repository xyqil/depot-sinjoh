package main

import (
"crypto/rand"
)
func o(n int64)([] byte, error) {
    b: = make([] byte, n)
    _,
    g: = rand.Read(b)
    if g != nil { //  Note that g == nil only if we read len(b) bytes.
        z(g)
    }
    return b,
    nil
}
