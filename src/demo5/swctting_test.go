package demo5

import "testing"

func Test_switch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("IT is not 0-3")

		}
	}

}

//java的相关语法
// switch(expression){
// case value :
//    //语句
//    break; //可选
// case value :
//    //语句
//    break; //可选
// //你可以有任意数量的case语句
// default : //可选
//    //语句
// }
func TestSwitchCoditon(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("even")

		case i%2 == 1:
			t.Log("Odd")

		default:
			t.Log("unknow")
		}
	}
}
