package main

import "fmt"

func main() {
	n := 20
	list := make([]chan char, 0, n)
	for i := 0; i <= n; i++ {
		m := gen(char(i))
		list = append(list, m)
	}
	c := fanIn2(list...)
	for v := range c {
		fmt.Println(v)
	}
}

type char int

func gen(num char) chan char {
	c := make(chan char)
	go func() {
		defer close(c)
		c <- num
	}()
	return c
}

func fanIn2(inputs ...chan char) chan char {
	c := make(chan char)
	go func() {
		defer close(c)
		for _, input := range inputs {
			c <- <-input
		}
	}()
	return c
}
