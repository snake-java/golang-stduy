package demo3

import "testing"

func Test_arryEqual(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 5, 6}
	d := [...]int{1, 2, 3, 4, 5, 7}
	t.Log(a == b)
	//长度不同的数据是无法比较，所有不能以a=c
	t.Log(d == c)
	t.Log("helloworld")
}

const (
	Open = 1 << iota
	Close
	Pending
)

func Test_byweiZero(t *testing.T) {
	a := 7
	//a &^ close 清掉close对应二进制位
	a = a &^ 3

	t.Log(a&Open == Open, a&Close == Close, a&Pending == Pending)

}
