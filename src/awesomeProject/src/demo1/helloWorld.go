package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("helloWorld")
	if len(os.Args) > 1 {
		fmt.Println("helloWorld", os.Args[0])
		fmt.Println("helloWorld", os.Args[1])
	}
	os.Exit(-1)
}
