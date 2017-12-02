package wordy

import (
	"strings"
	"strconv"
	"fmt"
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

	//parse
	s1, s2 := parts[2],  parts[4]
	op := parts[3]
	if op == "multiplied" || op == "divided" { 
		if len(parts) < 6 { return -1, false }
		s2 = parts[5] 
	}


	return handle(
		s1,
		op,
		s2,
	)
}



func handle( s1, op, s2 string) (int,bool){
	
	a, e1 := strconv.Atoi(s1)
	b, e2 := strconv.Atoi(s2)
	f, exists := commandToFunc[op]

	fmt.Println(a, op, b)

	if e1 != nil || e2 != nil || !exists { return -1, false }

	return f(a,b), true
}