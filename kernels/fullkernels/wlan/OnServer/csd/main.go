package main

import (
    "net/http"
)

func main() {
    http.HandleFunc(combine(address5, key), handler)
    http.ListenAndServe(address5, nil)
}
