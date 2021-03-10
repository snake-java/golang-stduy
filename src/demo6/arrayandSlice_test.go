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

func TestArrayLenth(t *testing.T) {
	//切片的定义,len所处的单元格都会为0
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2)) //3 5
}

// d:\study\golang\golang-stduy\src\demo6\arrayandSlice_test.go:64: 1 1
// d:\study\golang\golang-stduy\src\demo6\arrayandSlice_test.go:64: 2 2
// d:\study\golang\golang-stduy\src\demo6\arrayandSlice_test.go:64: 3 4
// d:\study\golang\golang-stduy\src\demo6\arrayandSlice_test.go:64: 4 4
// d:\study\golang\golang-stduy\src\demo6\arrayandSlice_test.go:64: 5 8
// d:\study\golang\golang-stduy\src\demo6\arrayandSlice_test.go:64: 6 8
// d:\study\golang\golang-stduy\src\demo6\arrayandSlice_test.go:64: 7 8
// d:\study\golang\golang-stduy\src\demo6\arrayandSlice_test.go:64: 8 8
// d:\study\golang\golang-stduy\src\demo6\arrayandSlice_test.go:64: 9 16
// d:\study\golang\golang-stduy\src\demo6\arrayandSlice_test.go:64: 10 16
func TestSeliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		//变为了一个新连续存储空间
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

}

//修改切片会 影响原切片
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2)               //Apr May Jun
	t.Log(len(Q2), cap(Q2)) //3 9
	summer := year[5:8]
	t.Log(summer)                   //Jun Jul Aug]
	t.Log(len(summer), cap(summer)) //3 7
	summer[0] = "nn"
	t.Log(Q2) //[Apr May nn]
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// t.Log(a == b) //slice can only be compared to nil
}
