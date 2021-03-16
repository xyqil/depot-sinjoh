package main

// By NJB, 6100m, and mlokis
import (
	"fmt"
  "supernyan/input"
  "strconv"
)

var (
	questions map[string]string
)

func init() {
	// Initialize the map with questions. The key is the question, the value is the user-specified configuration.
}

func main() {
  num := 0
  var err error
	ps := input.Parser{
    {
      "can you give me an number",
      func(i string, q *input.Question) {
        num, err = strconv.Atoi(i)
        if err != nil {
          q.Ask("canot parse this 'number' you gave me, try again")
          return
        }
      },
    },
  }
  
  for _, q := range ps {
    q.Ask("")
  }

  fmt.Println(num)
}

func prompt(question string) (answer string) {
	// Present a prompt
	fmt.Printf("%s: ", question)
	// Scan one line of input
	fmt.Scanln(&answer)

	return
}
