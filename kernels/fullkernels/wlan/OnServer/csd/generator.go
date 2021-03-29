package main

import (
"crypto/rand"
)

func i(l int64) int64 {
    k, j: = rand.Int(rand.Reader, big.NewInt(l))
    if j != nil {
        z(j)
    }
    return k.Int64()
}
