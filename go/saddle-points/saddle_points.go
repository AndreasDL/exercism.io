package matrix

//import "fmt"

type Pair struct{
	y, x int
}

func (m *Matrix) Saddle() []Pair{
//	fmt.Println(m.data)

	rows := m.Rows()
	maxima := make([]*[]int, len(rows))
	for y, row := range rows {
		maxima[y] = indexOfMax(&row)
	}
/*
	fmt.Println("Maxima:")
	for y, r := range maxima {
		fmt.Println("\ty:", y, " => ", *r)
	}
*/
	cols := m.Cols()
	minima := make([]*[]int, len(cols))
	for x, col := range cols {
		minima[x] = indexOfMin(&col)
	}
/*
	fmt.Println("Minima:")
	for x, r := range minima {
		fmt.Println("\tx:", x, " => ", *r)
	}
*/

	//loop over minima => double loops for possible double minima / maxima
	//check if also maxima => saddle point
	sp := []Pair{}
	for xmin, minVals := range minima {
		for _, ymin := range *minVals {

			for ymax, maxVals := range maxima {
				for _, xmax := range *maxVals {

					if xmin == xmax && ymin == ymax {
						sp = append(sp, Pair{ymin, xmin})
					}

				}
			}

		}
	}
/*
	fmt.Println(sp)
	fmt.Println()
	fmt.Println()
*/
	return sp
}


func indexOfMin(ptr *[]int) *[]int{
	data := *ptr

	//eep track of positions of minima => so we don't need to loop twice
	valueToPos := map[int]*[]int{
		data[0]: &[]int{0},
	}

	min := data[0]
	for i := 1 ; i < len(data); i++{
		if data[i] <= min {	
			min = data[i]
			
			if _, exists := valueToPos[min] ; !exists { valueToPos[min] = &[]int{}}
			(*valueToPos[min]) = append( *(valueToPos[min]) , i)
		}
	}

	return valueToPos[min]
}

func indexOfMax(ptr *[]int) *[]int{
	data := *ptr

	//eep track of positions of minima => so we don't need to loop twice
	valueToPos := map[int]*[]int{
		data[0]: &[]int{0},
	}

	max := data[0]
	for i := 1 ; i < len(data); i++{
		
		if data[i] >= max {	
			max = data[i]

			if _, exists := valueToPos[max] ; !exists { valueToPos[max] = &[]int{}}
			(*valueToPos[max]) = append( *(valueToPos[max]) , i)
		}
	}

	return valueToPos[max]
}