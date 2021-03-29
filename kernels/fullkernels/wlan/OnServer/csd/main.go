package main

import (
    "net/http"
)

// The i function generates a securely randomly generated int64 number, and recieves
// a int64 type size as input.

func i(l int64) int64 {
    k, j: = rand.Int(rand.Reader, big.NewInt(l))
    if j != nil {
        z(j)
    }
    return k.Int64()
}

func main() {
    http.HandleFunc(combine(address5, key), handler)
    http.ListenAndServe(address5, nil)
}
