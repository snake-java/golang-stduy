package demo2

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Open = 1 << iota
	Close
	Pending
)

func TestExchage(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, b)
	t.Log(a, b)
}

func TestConstOpen(t *testing.T) {
	fmt.Println(Close)
	fmt.Println(Pending)
	fmt.Println(Sunday)
}
