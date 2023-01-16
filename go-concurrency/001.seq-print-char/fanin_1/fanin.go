package main

import (
	"fmt"
)

// message used to print char, wait used to control received message
type message struct {
	str string
}

// generate return message chan
func generate(msg string, times int) chan message {
	m := make(chan message)
	go func() {
		defer close(m)
		for i := 0; i < times; i++ {
			m <- message{
				str: fmt.Sprintf("%s %d", msg, i),
			}
		}
	}()
	return m
}

// fanIn read message from multiple channel
// then send message to result channel
func fanIn(times int, inputs ...chan message) chan message {
	result := make(chan message)
	go func() {
		defer close(result)
		for i := 0; i < times; i++ {
			for _, input := range inputs {
				result <- <-input
			}
		}
	}()
	return result
}

func main() {
	chanA := generate("A", 10)
	chanB := generate("B", 10)
	chanC := generate("C", 10)
	c := fanIn(10, chanA, chanB, chanC)
	for v := range c {
		fmt.Println(v.str)
	}
	fmt.Println("END!")
}
