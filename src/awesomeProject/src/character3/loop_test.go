package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even", i)
		case 1, 3:
			t.Log("Odd", i)
		default:
			t.Log("it's not in 1-3", i)
		}
	}
}
func TestSwitchMultiCaseConditional(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even", i)
		case i%2 == 1:
			t.Log("Odd", i)
		default:
			t.Log("it's not in 1-3", i)
		}
	}
}
func Test_loop(t *testing.T) {
	//while循环
	n := 0
	for n <= 5 {
		fmt.Println("helloWorld", n)
		n++
	}

	//无线循环
	for {
		if n == 10 {
			break
		}
		fmt.Println("helloWorld", n)
		n++
	}
	//赋值做判断
	if c := 1; c > 0 {
		fmt.Println("helloWorld", c)
	}
	//switch 不需要break 不需要指定参数 可以指定支付串
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "liunx":
		fmt.Println("liunx X")
	default:
		fmt.Printf("%s.\n\r", os)
	}

	Num := 10
	switch {
	case 0 < Num && Num <= 3:
		fmt.Println("0-3")
	case 4 < Num && Num <= 10:
		fmt.Println("4-10")
	case 11 < Num:
		fmt.Println("11-无穷")
	}

}
