package day03

import (
	"fmt"
	"testing"
)

/*
  go 对隐士类型的转换的
*/
type Myint int64

func TestGoType(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c Myint
	// c = b 隐式类型转换不支持
	c = Myint(b)
	t.Log(b, a, c)
}

func TestPoint(t *testing.T) {
	var a int = 1
	aPtr := &a
	//不支持指针运算
	// aPtr = aPtr +1
	fmt.Printf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	//golang string的默认是空值 string默认是值类型，
	var c string
	fmt.Println("**", c, "**") // **  **
	if "" == c {
		fmt.Println("很对")
	}
}
