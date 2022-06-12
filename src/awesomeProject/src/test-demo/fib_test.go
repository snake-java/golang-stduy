package main

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
	Open    = 1 << iota //1
	Close               //2
	Pending             //4
)

/**
1 1 2 3 5 8
*/
func Test_Const(t *testing.T) {
	t.Log("Monday:", Monday)
	t.Log("Close:", Close)
	t.Log("Close:", Pending)
	a := 1
	t.Log(a&Open == Open, a&Close == Close, a&Pending == Pending)
}

func Test_fib(t *testing.T) {
	//var a int = 1
	//var b int = 1
	a := 1
	b := 1
	fmt.Print(a, " ")
	for i := 0; i < 10; i++ {
		fmt.Print(b, " ")
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()

}

func Test_Swap(t *testing.T) {
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp
	//go语言交换变量
	a, b = b, a
	t.Log(a, b)
}
