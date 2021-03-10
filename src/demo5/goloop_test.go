package demo5

import (
	"testing"
)

func Test_for(t *testing.T) {
	//    for{
	// 	   t.Log("1111")
	//    }
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
