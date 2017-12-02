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

	for i, row := range rows{
		r, l, err := parseRow(row)

		if err != nil {
			return nil, err
		}else if m.no_cols == 0 {
			m.no_cols = l
		} else if l != m.no_cols {
			return nil, errors.New("All rows must be of the same length")
		}
		
		m.data[i] = r
	}

	return &m, nil
}


//allocate empty [][]int with same dimensions as matrix
func allocate(height, width int) *[][]int{
	res := make([][]int, height)
	for i, _ := range res{
		res[i] = make([]int, width)
	}
	return &res
}

func (m Matrix) Rows() [][]int {
	rows := *allocate(m.no_rows, m.no_cols)

	//copy
	for i := 0; i < len(rows) ; i++ {
		for j := 0 ;  j < len(rows[i]); j++{
			rows[i][j] = m.data[i][j]
		}
	}

	return rows
}

func (m Matrix) Cols() [][]int{
	cols := *allocate(m.no_cols, m.no_rows)

	//copy
	for i := 0; i < len(cols) ; i++ {
		for j := 0 ;  j < len(cols[i]); j++{
			cols[i][j] = m.data[j][i]
		}
	}

	return cols
}

func (m Matrix) Set(row, col, val int) bool{
	
	if row >= m.no_rows || row < 0 ||
		col >= m.no_cols || col < 0 { 
			return false 
	}


	m.data[row][col] = val
	
	return true
}