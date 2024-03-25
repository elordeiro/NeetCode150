package utils

type Set map[byte]struct{}

func (s *Set) add(item byte) *Set {
	if s == nil {
		s := make(Set)
		s[item] = struct{}{}
		return &s
	}
	(*s)[item] = struct{}{}
	return s
}

func (s *Set) remove(item byte) {
	delete(*s, item)
}

func (s *Set) contains(item byte) bool {
	if s == nil {
		return false
	}
	_, found := (*s)[item]
	return found
}
