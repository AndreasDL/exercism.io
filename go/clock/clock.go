package clock

import "fmt"

type Clock struct{
	m int
}

var minutesPerDay int = 24*60

func New(h, m int) Clock {
	m = (h * 60 + m) % minutesPerDay
	if m < 0 {
		m += minutesPerDay
		m %= minutesPerDay
	}
	return Clock{m}
}

func (c Clock) Add(m int) Clock{
	return New(0 , c.m + m)
}


func (c Clock) String() string{
	h := int(c.m / 60)
	m := c.m % 60

	return fmt.Sprintf("%02d:%02d", h, m)
}