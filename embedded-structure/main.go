package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int

	// Element in embedded structure can be directly accessed, but can not be directly initialized.
	// Embedded structure can be accessed as a field, but no way to be initialized as a field.
	Contact
}

type Contact struct {
	Tel   string
	Email string
}

func main() {
	method1()

	method2()
}

func method1() {
	c := Contact{
		Tel:   "12345678",
		Email: "robin@xx.com",
	}

	// Error: values and fields are not match
	// p := Person{29, c}

	p := Person{"robin", 29, c}

	fmt.Println(p)
}

func method2() {
	// Error: mixture of field:value and value initializers.
	/*p := Person{
		Name: "robin",
		Age:  29,
		Contact{
			Tel:   "12345678",
			Email: "robin@xx.com",
		},
	}*/

	// Error: unknown field 'Tel' in struct literal of type Person.
	// Error: unknown field 'Email' in struct literal of type Person.
	/*p := Person{
		Name:  "robin",
		Age:   29,
		Tel:   "12345678",
		Email: "robin@xx.com",
	}*/

	// Error: too few values in struct initializer.
	// p := Person{"robin", 29}

	p := Person{
		Name: "robin",
		Age:  29,
	}
	p.Contact = Contact{
		Tel:   "12345678",
		Email: "robin@xx.com",
	}

	fmt.Println(p)

	// Element in embedded structure can be directly accessed
	fmt.Println(p.Tel)
}
