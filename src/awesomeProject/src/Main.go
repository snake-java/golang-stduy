package main

import (
	"fmt"
	"os"
)

func main() {
	fileOut, err := os.Open("D:\\study\\golang\\golang-stduy\\src\\awesomeProject\\src\\hello.txt")
	if err == nil {
		bytes := make([]byte, 1024)
		for {
			read, err := fileOut.Read(bytes)
			if err != nil {
				fmt.Println("Error reading file:", err)
				break
			}
			fmt.Println(string(read))
		}
	}

}
