package secret



var events []string

func init() {
	events = []string{
		"wink",
		"double blink"        ,
		"close your eyes"     ,
		"jump"                ,
		"Reverse"             ,
	}
}

func Handshake(code uint) []string{
	res := []string{}

	for i := 0 ; i < 4 ; i++ {

		remainder := code % 2
		if remainder == 1 {
			res = append(res, events[i])
		}

		code -= remainder
		code /= 2
	}

	//check reverse flag
	if code % 2 == 1 {
		reverseArrayInPlace(&res) //reverse
	}
	
	return res
}


func reverseArrayInPlace(ptr *[]string){
	s := *ptr
	for i, j := 0, len(s)-1 ; i < j ; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }

}
