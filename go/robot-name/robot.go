package robotname

import (
	"math/rand"
	"time"
	"strconv"
)


type Robot struct {
	name string
}

func Init(){
	rand.Seed(time.Now().UnixNano())	
}


const no_letters = int(2)
const no_digits  = int(3)
const letters    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//interesting article for optimization https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang


func(r *Robot) Name() string{
	if len(r.name) < no_letters + no_digits {
		r.Reset() //Pointer needed otherwise we are resetting the copy !!
	}

	return r.name
}

func (r *Robot) Reset() string{

	r.name = ""

	for i := 0; i < no_letters ; i++ {
		r.name += string( letters[ rand.Intn(len(letters)) ] )
	}

	for i := 0; i < no_digits ; i++  {
		r.name += strconv.Itoa( rand.Intn(9) )
	}

	return r.name
}