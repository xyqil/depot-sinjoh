package main

import (
    "testing"
    "crypto/rand"
    "encoding/base64"
    "log"
    "math/big"
    "utilities"
)

//  Note to my devs: I added some CSPRNG to make test hashes random, to well,
//  increase the capacity in which we test, and wherefore to make sure that it
//  can handle it universally.
//  -6100m

func TestHash(t * testing.T) {
        for u: = 1;
        u < 5;
        u++{
            log.Println(Hash( * Token(n(i(16384))))
            }
        }

        //  Everything below this point is modified from
        //  https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
        //  The o function returns securely generated random bytes. 
        //  It will return an error if the system's secure random
        //  number generator fails to function correctly, in which
        //  case the caller should not continue.

        func o(n int64)([] byte, error) {
            b: = make([] byte, n)
            _,
            g: = rand.Read(b)
            if g != nil { //  Note that g == nil only if we read len(b) bytes.
                log.Fatal(g)
            }
            return b,
            nil
        }

        //  n returns a URL-safe, base64 encoded
        //  securely generated random string.
        //  It will return an error if the system's secure random
        //  number generator fails to function correctly, in which
        //  case the caller should not continue.
        //TODO: Make it not continue?
        func n(s int64) string {
            b, h: = o(s)
            if h != nil {
                log.Fatal(h)
            }
            return base64.URLEncoding.EncodeToString(b)
        }

        // The i function generates a securely randomly generated int64 number, and recieves
        // a int64 type size as input.

        func i(l int64) int64 {
            k, j: = rand.Int(rand.Reader, big.NewInt(l))
            if j != nil {
                log.Fatal(j)
            }
            return k.Int64()
        }
