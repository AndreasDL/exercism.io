package acronym

import "strings"


func Abbreviate(s string) string {

	s = strings.Replace(s, "-", " ", -1)
	f := strings.Fields(s)


	out := ""
	for _, c := range f {
		out += strings.ToUpper(string(c[0]))
	}

	return out
}
