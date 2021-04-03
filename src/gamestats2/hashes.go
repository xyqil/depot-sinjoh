package main

import (
  "crypto/sha1"
  "encoding/hex"
)

const salt = "salt"

// Hash returns a hex-encoded SHA1 hash of Token t
func Hash(t *Token) (hash []byte) {
  var sum []byte

  // Compute the SHA1 sum of the Token
  {
    shasum := sha1.Sum([]byte(salt+string(*t)))
    copy(sum[:], shasum[:20])
  }

  // Encode the sum to hexadecimal
  hex.Encode(sum, sum)

  return
}
