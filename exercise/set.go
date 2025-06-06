package exercise

type Set map[string]struct{}

func New() Set {
	return map[string]struct{}{}
}

func NewFromSlice(l []string) Set {
	newSet := make(map[string]struct{})
	for _, d := range l {
		newSet[d] = struct{}{}
	}
	return newSet
}

func (s Set) String() string {
	var result string
	result = `{`
	for key, _ := range s {
		result += `"` + key + `", `
	}
	if len(s) > 0 {
		result = result[:len(result)-2]
	}
	return result + `}`
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	for ele := range s1 {
		if _, ok := s2[ele]; !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return true
	}
	for ele := range s1 {
		if _, ok := s2[ele]; ok {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	if len(s1) == 0 || len(s2) == 0 {
		return New()
	}
	result := New()
	for ele := range s1 {
		if _, ok := s2[ele]; ok {
			result.Add(ele)
		}
	}
	return result
}

func Difference(s1, s2 Set) Set {
	if len(s1) == 0 || len(s2) == 0 {
		return s1
	}
	diff := New()
	for ele := range s1 {
		if _, ok := s2[ele]; !ok {
			diff.Add(ele)
		}
	}
	return diff
}

func Union(s1, s2 Set) Set {
	for elem := range s2 {
		s1.Add(elem)
	}
	return s1
}
