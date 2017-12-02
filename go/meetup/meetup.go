package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota 
	Second
	Third
	Fourth
	Last
	Teenth
)


func Day(sched WeekSchedule, weekDay time.Weekday, month time.Month, year int) int{

	//start day, depending on schedule, month & leap year
	day := 1
	switch sched {
		case First : day = 1
		case Second: day = 8
		case Third : day = 15
		case Fourth: day = 22
		case Teenth: day = 13
		case Last  :
			switch month {
				case time.January, time.March,   time.May, time.July, 
					 time.August,  time.October, time.December: 
					day = 31 - 6
				case time.April, time.June, time.September, time.November: 
					day = 30 - 6
				case time.February:
					if (year % 100 != 0 && year % 4 == 0) || year % 400 == 0 { //leap year
						day = 29 - 6
					} else { //non-leap year
						day = 28 - 6
					}
			}
	}

	return firstWeekdayAfter(
					weekDay,
					time.Date(year, month, day, 0,0,0,0,time.UTC),
		   )
}

//add days till we have the right weekday !
func firstWeekdayAfter(weekDay time.Weekday, date time.Time) int{
	for weekDay != date.Weekday() { date = date.AddDate(0,0, 1) }

	return date.Day()
}