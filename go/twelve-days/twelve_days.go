package twelve 


const no_verses = int(12)
var days []string = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var items []string = []string{
	"a Partridge in a Pear Tree", 
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Verse(day int) string{
	day-- //array indexing starts at ZERO

	s := "On the "
	s += days[day]
	s += " day of Christmas my true love gave to me, "

	for i:= day ; i > 0; i-- {
		s += items[i] + ", "
	}

	//last one with and
	if day > 0 {
		s += "and "
	}
	s += items[0] + "."

	return s
}

func Song() string {
	s := ""
	for i := 1 ; i <= no_verses; i++{
		s += Verse(i) + "\n"
	}

	return s
}

