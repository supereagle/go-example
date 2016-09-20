package main

import (
	"fmt"
)

func main() {
	err := AppendToFile("test.txt", "Hello, Go!\n")
	if err != nil {
		fmt.Println("Fail to append the content to file!")
	}
}
