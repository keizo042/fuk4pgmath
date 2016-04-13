package main

import (
	"fmt"
	"strconv"
)

type N (chan interface{})
type Z int

func Zero() N {
	n := make(N, 1)
	n <- Z(0)
	return n
}

func (n N) IsZero() bool {
	m := <-n
	switch m.(type) {
	case Z:
		n <- m
		return true
	default:
		n <- m
		return false

	}
}

func (n N) Succ() N {
	m := make(N, 1)
	m <- n
	return m
}

func (n N) Int() int {
	i := 0
	m := n
	for {
		l := <-m
		switch l.(type) {
		case N:
			m = l.(N)
			i++
		case Z:
			return i
		default:
			return -1
		}
	}
}

func (n N) String() string {
	return strconv.Itoa(n.Int())

}

func main() {
	n := Zero()
	n = n.Succ()
	n = n.Succ()
	n = n.Succ()
	n = n.Succ()
	fmt.Println(n)

}
