package main

import (
	"fmt"

	"github.com/supereagle/go-example/lint/students"
)

func main() {
	s := students.New("Robin", 30)

	// Output student information.
	fmt.Printf("Student: name: %s, age: %d\n", s.GetName(), s.GetAge())
}
