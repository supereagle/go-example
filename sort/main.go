package main

import (
	"sort"

	"github.com/caicloud/pkg/log"
)

type Student struct {
	name string
	age  int
}

// Students represents a set of students, and support to sort their order.
type Students []Student

func (s Students) Len() int { return len(s) }
func (s Students) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Ensure Students implements sort.Interface.
var _ sort.Interface = (*Students)(nil)

func main() {
	ss := Students{
		Student{
			name: "Jack",
			age:  20,
		},
		Student{
			name: "Rose",
			age:  18,
		},
	}

	log.Infof("Students before sorted: %+v", ss)
	sort.Sort(ss)
	log.Infof("Students after sorted: %+v", ss)
}
