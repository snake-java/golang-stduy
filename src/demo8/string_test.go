package demo8

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	fmt.Println(s)
	s = "hello"
	fmt.Println(len(s))
	s = "中"
	c := []rune(s)
	fmt.Print(len(c))
	fmt.Println(c)
	t.Logf("中的 unicode %x", c[0]) //4e 2d
	t.Logf("中 utf-8 %x", s)       //e4 b8 ad
}

func TestLoopString(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		fmt.Println(c)
		//%c编码 %d 数字
		fmt.Println(string(c))
		t.Logf("%[1]c %[1]d", c)
	}
}

func TestStringFunction(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Logf(part)
	}

}

func TestStringJoin(t *testing.T) {
	s := "a,b,c"
	s2 := strings.Split(s, ",")
	s1 := strings.Join(s2, "-")
	t.Log(s1)
	s3 := strconv.Itoa(10)
	fmt.Printf("%T\n", s3)
	// strconv.Atoi()
	if s4, de := strconv.Atoi(s3); de == nil {
		fmt.Println(s4 + 10)
	}

}
