package utils

type Set struct{
	contents map[string]bool
}

func (set Set) Contains(element string) bool  {
	_, ok := set.contents[element]
	return ok
}

func (set Set) add(element string) {
	set.contents[element] = true
}

func NewSet(elements... string) Set{
	s := Set{ map[string]bool{} }
	for _, e := range elements {
		s.add(e)
	}
	return s
}
