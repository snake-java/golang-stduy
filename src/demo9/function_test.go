package demo10

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func retrunvalues() (int, int) {
	return rand.Intn(10), rand.Intn(10)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}
func slowFunc(i int) int {
	time.Sleep(time.Second * 2)
	return i
}

func changeArgs(args ...int) int {
	if len(args) == 1 {
		return args[0]
	}
	sum := 0
	for _, a := range args {
		sum += a
	}
	return sum

}
func TestTimeSpent(t *testing.T) {
	tsf := timeSpent(slowFunc)
	t.Log(tsf(10))
}

func TestFunction(t *testing.T) {
	v1, v2 := retrunvalues()
	t.Log(v2)
	t.Log(v1)
}

func TestChangeArgs(t *testing.T) {
	t.Log(changeArgs(1, 2, 3, 4, 5))
}

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Print("你好北京！")
	}()
	fmt.Println("你好广州！")
	//中断程序
	panic("fatal error")
}
