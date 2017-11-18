package series

//import "fmt"

func All(n int, s string) []string {
	res := make([]string, len(s) - n + 1) //allocate correct size

	for i := 0 ; i <= len(s) - n; i++{
		res[i] = s[i:i+n]
	}

	return res
}

func UnsafeFirst(n int, s string) string{
	return s[:n]
}

func First(n int, s string) (string, bool){
	if len(s) < n {
		return s, false
	}

	return s[:n], true
}