package sets

type StringSets map[string]struct{}

func NewStringSets(elements []string) StringSets {
	s := make(map[string]struct{})

	for _, element := range elements {
		s[element] = struct{}{}
	}

	return s
}

func (s StringSets) Contains(element string) bool {
	_, found := s[element]
	return found
}
