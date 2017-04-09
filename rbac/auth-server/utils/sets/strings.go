package sets

type Empty struct{}

// StringSet Set of strings, implemented via map[string]struct{} for minimal memory consumption.
type StringSet map[string]Empty

func NewStringSet(items []string) StringSet {
	s := StringSet{}

	for _, item := range items {
		s[item] = Empty{}
	}

	return s
}

// Has returns true if and only if item is contained in the set.
func (s StringSet) Has(item string) bool {
	_, found := s[item]
	return found
}

// HasAll returns true if and only if all items are contained in the set.
func (s StringSet) HasAll(items ...string) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s StringSet) HasAny(items ...string) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}

// Delete removes the item from the set.
func (s StringSet) Delete(item string) {
	delete(s, item)
}
