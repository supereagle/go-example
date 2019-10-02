package students

type Students struct {
	name string
	age  int
}

func New(name string, age int) *Students {
	return &Students{
		name: name,
		age:  age,
	}
}

func (s *Students) GetName() string {
	return s.name
}

func (s *Students) IsYoung() bool {
	return s.age < 18
}

func (s *Students) GetAge() int {
	return s.age
}

type Score struct {
	cource string
	score  int
}

func NewScore(cource string, score int) *Score {
	return &Score{
		cource: cource,
		score:  score,
	}
}
