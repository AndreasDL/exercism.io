package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {



	return t.Add(time.Duration(1e9) * time.Second)
}
