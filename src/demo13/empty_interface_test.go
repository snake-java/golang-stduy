package demo13

import "fmt"

func DoSomething(p interface{}) {
	// if i, ok := p.(string); ok {
	// 	fmt.Println("is string", i)
	// 	return
	// }
	// if i, ok := p.(int); ok {
	// 	fmt.Println("is string", i)
	// 	return
	// }
	// fmt.Println("i don't type")
	switch v := p.(type) {
	case int:
		fmt.Println("is int", v)
	case string:
		fmt.Println("is string", v)
	default:
		fmt.Println("i do know ")
	}
}
