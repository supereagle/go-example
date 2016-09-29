package main

import (
	"fmt"
)

type Student struct {
	Name    string
	Age     int
	Classes []string
}

func main() {
	student := Student{
		Name:    "John",
		Age:     25,
		Classes: []string{"Chinese", "English", "Computer"},
	}

	// Marshal Json object to string
	stdStr, err := Marshal2JsonStr(student)
	if err != nil {
		fmt.Printf("Fail to marshal json object as %s", err.Error())
		return
	}
	fmt.Println(stdStr)

	// Unmarshal string to Json object
	var stdObj Student
	err = UnmarshalJsonStr2Obj(stdStr, &stdObj)
	if err != nil {
		fmt.Printf("Fail to unmarshal string to json object as %s", err.Error())
		return
	}
	fmt.Printf("The object is: %v", stdObj)
}
