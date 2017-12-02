package binarysearch

//import "fmt"

func SearchInts(slice []int, key int) int {
//	fmt.Println(slice, key)

	start, stop := 0, len(slice)
	pos := int((start + stop) / 2)
//	fmt.Println("\t", start, stop, pos)
	for start < stop && slice[pos] != key {

		if slice[pos] < key { // go right
			start = pos + 1
		} else { //go left
			stop  = pos - 1
		}

		pos = int((start + stop) / 2)
//		fmt.Println("\t", start, stop, pos)
	}

	if pos < len(slice) && slice[pos] == key { return pos }

	return -1
}