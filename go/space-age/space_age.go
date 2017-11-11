package space

import "math"


type Planet string

var orbital_period_on_earth = 31557600.0

var planet_to_orbital_period_seconds = map[Planet]float64 {
	"Earth"   : orbital_period_on_earth,
	"Mercury" : 0.2408467  * orbital_period_on_earth,
    "Venus"   : 0.61519726 * orbital_period_on_earth,
    "Mars"    : 1.8808158  * orbital_period_on_earth,
    "Jupiter" : 11.862615  * orbital_period_on_earth,
    "Saturn"  : 29.447498  * orbital_period_on_earth,
    "Uranus"  : 84.016846  * orbital_period_on_earth,
    "Neptune" : 164.79132  * orbital_period_on_earth,
}

func round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}


func Age(s float64, p Planet) float64{

	return round(s / planet_to_orbital_period_seconds[p],.5,2)

}