package main

import (
"log"
)
func z(e error) {
    if e != nil {
        log.Fatal(e)
    }
}
