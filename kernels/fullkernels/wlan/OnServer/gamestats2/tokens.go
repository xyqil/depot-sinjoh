package main

import (
    "time"
    "math/rand"
)

// Runes out of which Tokens are generated
var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
  // Ensure accurate pseudo-random number generation
  rand.Seed(time.Now().UnixNano())
}

// Token is a random 32-character long string
type Token string

// GenerateToken generates a random Token.
func GenerateToken() Token {
  buf := make([]rune, 32)
  for i := range buf {
    buf[i] = runes[rand.Intn(len(runes))]
  }

  return Token(string(buf))
}
