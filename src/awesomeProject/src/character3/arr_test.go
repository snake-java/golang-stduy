package main

import (
	"testing"
)

const (
	Open    = 1 << iota //1
	Close               //2
	Pending             //4
)

/**
1 1 2 3 5 8
1&^0 ---1
1&^1 ----0
0&^1 ----0
0&^0 ----0
*/
func Test_BinaryClear(t *testing.T) {
	a := 7
	a = a &^ Open
	t.Log(a&Open == Open, a&Close == Close, a&Pending == Pending)
}

func Test_arr(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	arr1 := [...]int{1, 2, 3, 6}
	//arr2 := [...]int{1, 2, 3, 6, 4}
	arr3 := [...]int{1, 2, 3, 4}
	t.Log(arr == arr1)
	t.Log(arr == arr3)
	//invalid operation: arr == arr2 (mismatched types [4]int and [5]int)
	//t.Log(arr == arr2)
}
