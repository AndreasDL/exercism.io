package minesweeper

import (
	"bytes"
	"strconv"
	"errors"
)

type Board [][]byte

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}

func (b *Board) Count() error{
	if !b.check() { return errors.New("Useful error Message") }

	for l, line := range *b {
		for c, place := range line {

			if place == byte(' ') {
				if val := b.getSurrBombs(l,c) ; val != byte('0'){ 
					(*b)[l][c] = val
				}
			}
		}
	}

	return nil
}

func (ptr *Board) check() bool{
	b := *ptr
	if len(b) == 0 { return false }

	dim := len(b[0])
	for i, line := range b {

		if len(line) != dim { 
			return false 
		} else if (i == 0 || i == len(b)-1) && (line[0] != byte('+') || line[len(line)-1] != byte('+')){ //first & last line
			return false
		} else if (i > 0 && i < len(b)-1) {//lines in between

			if line[0] != byte('|') || line[dim-1] != byte('|') {	
				return false
			}

			//check chars
			for _, c := range line[1:len(line)-1]{
				if c != byte(' ') && c != byte('*') { 
					return false 
				}
			}
		}
	}

	return true
}

func (ptr *Board) getSurrBombs(l,c int) byte{
	b := *ptr

	dirs := []struct{
		dx, dy int
	}{
		{-1,-1}, {-1, 0}, {-1, 1},
		{ 0,-1}, { 0, 0}, { 0, 1},
		{ 1,-1}, { 1, 0}, { 1, 1},
	}

	res := 0
	for i := 0; i < len(dirs) ; i++ {
		dx := dirs[i].dx + c
		dy := dirs[i].dy + l

		if dx < 0 || dx > len(b[0]) { 
			continue
		} else if dy < 0 || dy > len(b){
			continue
		}

		if b[dy][dx] == byte('*'){
			res++
		}

	}

	return byte( strconv.Itoa(res)[0] )
}