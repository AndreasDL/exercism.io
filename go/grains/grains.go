package grains

import "errors"
/*
Calculate the number of grains of wheat on a chessboard given that the number
on each square doubles.

There once was a wise servant who saved the life of a prince. The king
promised to pay whatever the servant could dream up. Knowing that the
king loved chess, the servant told the king he would like to have grains
of wheat. One grain on the first square of a chess board. Two grains on
the next. Four on the third, and so on.

There are 64 squares on a chessboard.

Write code that shows:
- how many grains were on each square, and
- the total number of grains
*/

func Square(i int) (uint64, error) {

	if i < 1 || i > 64{
		return 0, errors.New("input number is not valid; there are between 1 and 64 pieces on the board.")
	}

	return uint64(1) << uint(i-1), nil
}

func Total() uint64 {

	sum := uint64(0)

	for i := 0; i < 64; i++ {
		sum += 1 << uint(i)
	}

	return sum
}