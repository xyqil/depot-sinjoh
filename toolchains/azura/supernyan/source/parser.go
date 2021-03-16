package input

import (
  "fmt"
)

// Parser stores questions and runs them
type Parser []Question

// Question ...
type Question struct {
  Message string
  Handler func(input string, q *Question)
}

// Ask asks for an answer an runs handler on user input
func (q *Question) Ask(message string) error {
  if message == "" {
    message = q.Message
  }
  fmt.Printf("%s :", message)
  ans := ""
  _, err := fmt.Scanln(&ans)
  if err != nil {
    return err
  }
  q.Handler(ans, q)
  return nil
} 