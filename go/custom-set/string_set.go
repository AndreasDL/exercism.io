package stringset

//import "fmt"

type Set map[string]bool


func New() Set{
	return Set{}
}

func NewFromSlice(slice []string) Set{

	set := Set{}

	for _, str := range slice {
		set[str] = true
	}

	return set
}


func (s Set) String() string{
	str := "{"

	for k, _ := range s {
		str += "\"" + k + "\", "
	}

	if len(str) > 2 {
		str = str[:len(str)-2]	
	}
	
	str += "}"

	return str
}

func (s Set) IsEmpty() bool{
	return len(s) == 0
}

func (s Set) Has(str string) bool{
	_, exists := s[str]

	return exists
}

func (s *Set) Add(str string) {
	(*s)[str] = true
}

func Subset(s1, s2 Set) bool{

	for k, _ := range s1 {
		if _, exists := s2[k] ; !exists { return false }
	}

	return true
}

func Disjoint(s1, s2 Set) bool{

	for k, _ := range s1 {
		if _, exists := s2[k] ; exists { return false }
	}

	return true
}

func Equal(s1, s2 Set) bool{

	return len(s1) == len(s2) && Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set{

	intersect := Set{}

	for k , _ := range s1 {
		if _, exists := s2[k] ; exists { intersect[k] = true }
	}
	return intersect
}

func Difference(s1, s2 Set) Set{ //only in one direction
	diff := Set{}

	for k , _ := range s1 {
		if _, exists := s2[k] ; !exists { diff[k] = true }
	}
/*
	for k, _ := range s2 {
		if _, exists := s1[k] ; !exists { diff[k] = true }
	}
*/

	return diff
}

func Union(s1, s2 Set) Set{
	
	union := Set{}

	for k, _ := range s1 { union[k] = true }
	for k, _ := range s2 { union[k] = true }

	return union
}