package foodchain


type item struct {
	animal string
	reaction string
	fullDesc string
}

var items = []item{
	item{
		animal   : "fly",
		fullDesc : "fly",
		reaction : "I don't know why she swallowed the fly.",
	},
	item{
		animal   : "spider",
		fullDesc : "spider that wriggled and jiggled and tickled inside her",
		reaction : "It wriggled and jiggled and tickled inside her.",
	},
	item{
		animal   : "bird",
		fullDesc : "bird",
		reaction : "How absurd to swallow a bird!",
	},
	item{
		animal   : "cat",
		fullDesc : "cat",
		reaction : "Imagine that, to swallow a cat!",
	},
	item{
		animal   : "dog",
		fullDesc : "dog",
		reaction : "What a hog, to swallow a dog!",
	},
	item{
		animal   : "goat",
		fullDesc : "goat",
		reaction : "Just opened her throat and swallowed a goat!",
	},
	item{
		animal   : "cow",
		fullDesc : "cow",
		reaction : "I don't know how she swallowed a cow!",
	},
	item{
		animal   : "horse",
		fullDesc : "horse",
		reaction : "She's dead, of course!",
	},
}


func Verse(v int) string{
	v-- //array indexing starts at 0
	if v < 0 || v > len(items) {
		return ""
	}

	verse := "I know an old lady who swallowed a " + items[v].animal + ".\n"
	
	//exceptions
	if v > 0 { //no reaction for first item
		verse += items[v].reaction
	}

	if v == 7 { //last one is different
		return verse
	} else if v > 0 { //put the newline for all verses between 1 & 7
		verse += "\n"
	}
	

	for i := v ; i > 0 ; i-- {
		verse += "She swallowed the " + items[i].animal + " to catch the " + items[i-1].fullDesc + ".\n"
	}
	verse += items[0].reaction + " Perhaps she'll die."

	return verse
}

func Verses(start, stop int) string{
	s := Verse(start)

	for i := start + 1 ; i <= stop ; i++{
		s += "\n\n"
		s += Verse(i)
	}
	return s
}

func Song() string{
	return Verses(1,8)
}