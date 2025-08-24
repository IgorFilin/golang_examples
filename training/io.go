package training

import (
	"fmt"
	"io"
)

type TReader string

func (numbers TReader) Read(n []byte) (int, error) {
	count := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			count++
			n[i] = numbers[i]
		}
	}
	return count, io.EOF
}

func IoTraining() {
  var counters TReader = "123456789"

  var buffer []byte = make([]byte, len(counters))

  counters.Read(buffer)

  fmt.Println(buffer)
  fmt.Println(string(buffer))
}