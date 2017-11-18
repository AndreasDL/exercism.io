package house

type item struct{
	item, verb string
}

var items []item

func init(){
	items = []item{
		item{
			item: "malt",
			verb: "ate",
		},item{
			item: "rat",
			verb: "killed",
		},item{
			item: "cat",
			verb: "worried",
		},item{
			item: "dog",
			verb: "tossed",
		},item{
			item: "cow with the crumpled horn",
			verb: "milked",
		},item{
			item: "maiden all forlorn",
			verb: "kissed",
		},item{
			item: "man all tattered and torn",
			verb: "married",
		},item{
			item: "priest all shaven and shorn",
			verb: "woke",
		},item{
			item: "rooster that crowed in the morn",
			verb: "kept",
		},item{
			item: "farmer sowing his corn",
			verb: "belonged to",
		},item{
			item: "horse and the hound and the horn",
			verb: "",
		},
	}

}

func Verse(paragraph int) string{
	if paragraph == 1 {
		return "This is the house that Jack built."
	}

	paragraph -= 2
	
	song := "This is the " + items[paragraph].item + "\n"
	for i := paragraph -1 ; i >= 0 ; i-- {
		song += "that " + items[i].verb + " the " + items[i].item + "\n"
	}
	song += "that lay in the house that Jack built."

	return song
}

func Song() string{
	
	song := Verse(1)
	
	for paragraph := 2 ; paragraph < len(items) +2 ; paragraph++ {
		
		song += "\n\n"
		song += Verse(paragraph)

	}
	return song
}