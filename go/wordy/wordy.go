package wordy

import (
	"strings"
	"strconv"
)

var commandToFunc = map[string]func(int,int) int{
	"plus"       : func(a,b int) int {return a + b},
	"minus"      : func(a,b int) int {return a - b},
	"multiplied" : func(a,b int) int {return a * b},
	"divided"    : func(a,b int) int {return a / b},
}


func Answer(s string) (int, bool){
	//clean input
	if len(s) < 1 { return -1, false}
	s = s[:len(s)-1] //remove trailing questionmark

	//split line
	parts := strings.Split(s, " ")
	if len(parts) < 5 { return -1, false }

	//parse first value
	res, e := strconv.Atoi(parts[2])
	if e != nil { return -1, false }

	//loop over operator, value pairs !
	ok := true
	for i := 3 ; i < len(parts) ; i+= 2{
		op := parts[i]
		if op == "multiplied" || op == "divided" { i++ } //"divided by, multiplied by" 

		val := parts[i+1]
		
		res, ok = handle(res, op, val)
		if !ok { return -1, false }
	}

	return res, true
}



func handle( a int, op, s2 string) (int,bool){	
	b, e := strconv.Atoi(s2)
	f, exists := commandToFunc[op]

	if e != nil || !exists { return -1, false }

	return f(a,b), true
}