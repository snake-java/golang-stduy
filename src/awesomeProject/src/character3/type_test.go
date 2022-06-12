package main

import (
	"testing"
)

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64 = 1
	//cannot use a (type int32) as type int64 in assignment
	//b = a
	//显示类型转化
	b = int64(a)
	t.Log(a, b)
}

//Invalid operation: aptr + 1 (mismatched types *int and untyped int)
func TestPoint(t *testing.T) {
	//a := 1
	//aptr := &a
	//apte = aptr + 1 // 不支持

}

//字符串 默认是空字符串
func TestString(t *testing.T) {
	var s string
	if s == "" {
		t.Log("helloWorld")
	}
	t.Log(len(s))

}
