package leap

/*
on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    unless the year is also evenly divisible by 400
*/


func IsLeapYear(y int) bool {

	divBy4 :=  ( (y % 4) == 0 )
	divBy100 := ( (y % 100) == 0 )
	divBy400 := ( (y % 400) == 0 )

	return divBy4 && (!divBy100 || (divBy100 && divBy400) )
}
