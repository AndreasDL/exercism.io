package school

import "sort"


type Grade struct{
	year int
	names []string
}

type School struct{
	enrollment map[int]*Grade
}

func New() *School{
	return &School{
		enrollment: map[int]*Grade{},
	}
}

func (s *School) Add(name string, grade int){

	g, ex := s.enrollment[grade]

	if !ex { //create if no exists
		g = &Grade{grade, []string{}} 
		s.enrollment[grade] = g
	}

	g.names = append(g.names, name)
}

func (s *School) Grade(grade int) []string{

	g, ex := s.enrollment[grade]

	if !ex { return []string{} }

	sort.Strings(g.names)
	return g.names
}

func(s *School) Enrollment() []Grade{
	
	res := []Grade{}

	for _, v := range s.enrollment {

		sort.Strings(v.names)

		res = append(res, *v)
	}

	sort.Slice(res, 
		func(i,j int)bool { return res[i].year < res[j].year },
	)

	return res
} 