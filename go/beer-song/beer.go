package beer

import (
	"strconv"
	"errors"
	//"fmt"
)

func Song() string{
	song, _ := Verses(99, 0)
	return song
}

func Verses(start, stop int) (string, error){
	if start < stop { return "", errors.New("start should be >= stop") }

	res := ""

	for i := start ; i >= stop ; i-- {
		v, e := Verse(i)

		if e != nil { return "", e }

		res += v
		res += "\n"
	}

	return res, nil
}

func Verse(v int) (string, error) {
	//lazy I know, but tbh i don't like the song exercises.
	if v < 0 || v > 100 {
		return "", errors.New("Verse out of range")
	}else if v == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\n" +
			"Go to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}else if v == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\n" +
			"Take it down and pass it around, no more bottles of beer on the wall.\n", nil
	}else if v == 2 {
		return "2 bottles of beer on the wall, 2 bottles of beer.\n" +
			"Take one down and pass it around, 1 bottle of beer on the wall.\n", nil
	}

	bottles := strconv.Itoa(v) + " bottles of beer"
	res := bottles
	res += " on the wall, "
	res += bottles
	res += ".\n"

	bottles = strconv.Itoa(v-1) + " bottles of beer"
	res += "Take one down and pass it around, "
	res += bottles
	res += " on the wall.\n"

	return res, nil
}