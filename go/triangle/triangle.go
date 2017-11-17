package triangle

import (
	"sort"
	"math"
)

type Kind int

const (
    NaT = iota // not a triangle
    Equ = iota // equilateral
    Iso = iota // isosceles
    Sca = iota // scalene
    Deg = iota // degenerate (dig deeper exercise)
)


func KindFromSides(a, b, c float64) Kind {

	sides := []float64{a,b,c}
	sort.Float64s(sides)

	if sides[0] <= 0 || math.IsNaN(sides[0]) || math.IsInf(sides[0], -1) || math.IsInf(sides[2], +1)  { //sides with zero or below zero length or NaN
		return NaT
	} else if sides[0] + sides[1] < sides[2] {
		return NaT
/*
	Dig deeper exercise
	} else if sides[0] + sides[1] == sides[2] {
		return Deg
*/
	} else if sides[0] == sides[1] && sides[0] == sides[2] {
		return Equ
	} else if sides[0] == sides[1] || sides[1] == sides[2] {
		return Iso
	} else {
		return Sca
	}
}
