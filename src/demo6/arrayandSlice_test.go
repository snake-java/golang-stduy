package demo6

import (
	"fmt"
	"testing"
)

func TestArrayIntit(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1, 3, 3}
	arr2 := [...]int{1, 3, 4}
	t.Log(arr[1], arr[2])
	t.Log(arr1[1])
	t.Log(arr2[1])

}

func TestArrayFor(t *testing.T) {
	arr2 := [...]int{1, 3, 4}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}
	for _, e := range arr2 {
		t.Log(e)
	}
}

func TestArrayTwo(t *testing.T) {
	a := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}} /* 第三行索引为 2 */
	/* 输出数组元素 */
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

// :47: [3]
// :48: [1 3]
// :49: [3 4 5 7]
// :50: [4 5 7]
func TestSlice(t *testing.T) {
	a := [...]int{1, 3, 4, 5, 7}
	t.Log(a[1:2]) //不包含结束位置
	t.Log(a[:2])
	t.Log(a[1:])
	t.Log(a[2:len(a)])

}
