package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct{
	no_rows, no_cols int
	data [][]int
}

//returns row []int, lenth of row int, and error
func parseRow(s string) ([]int, int, error){ //array , len
	
	s = strings.Trim(s, " ")
	srow := strings.Split(s, " ")
	irow := make( []int, len(srow) )
	
	for i, v := range srow {
		n, err := strconv.Atoi(v)
		irow[i] = n
		if err != nil {
			return nil, 0, errors.New("Int overflow " + srow[i] + " becomes" + strconv.Itoa(irow[i]))
		}
	}

	return irow, len(irow), nil
}

func New(s string) (*Matrix, error){
	s = strings.Trim(s, " ")
	rows := strings.Split(s, "\n")

	m := Matrix{
		data: make([][]int, len(rows)),
		no_rows: len(rows),
	}

	for _, row := range rows{
		r, l, err := parseRow(row)

		if err != nil {
			return nil, err
		}else if m.no_cols == 0 {
			m.no_cols = l
		} else if l != m.no_cols {
			return nil, errors.New("All rows must be of the same length")
		}
		
		m.data = append(m.data, r)
	}

	return &m, nil
}




func (m Matrix) Rows() []int {
	return nil
}

func (m Matrix) Cols() []int{
	return nil
}

func (m Matrix) Set() []int{
	return nil
}