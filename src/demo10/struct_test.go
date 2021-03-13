package demo10

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Name string
	Age  uint
	Id   string
}

func (e Employee) SayName1() {
	fmt.Printf("address is %x", unsafe.Pointer(&e.Name))
	fmt.Println(e.Name)
}
func (e *Employee) SayName() {
	fmt.Printf("address is %x", unsafe.Pointer(&e.Name))
	fmt.Println(e.Name)
}

func Test_employee(t *testing.T) {
	//指针和实例进行访问对象的区别
	e := Employee{"zhangsan", 29, "1"}
	e1 := Employee{Name: "lisi", Age: 28, Id: "2"}
	e2 := new(Employee)
	e.SayName1()
	e.SayName1()
	//address is c000078540zhangsan  使用实例的传递是会产生对象的复制
	// address is c000078570zhangsan
	e2.Id = "3"
	e2.Name = "wangwu"
	e2.SayName()
	e2.SayName()
	// address is c000078510wangwu 使用引用的传递不会有对象的复制
	// address is c000078510wangwu
	t.Log(e)

	t.Log(e1)
	t.Log(e2)
}
