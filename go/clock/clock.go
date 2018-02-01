package clock

import "fmt"

type Clock struct {
	min int
}

func New(hour, min int) Clock {
	totalMin := (hour*60 + min) % 1440
	if totalMin < 0 {
		for oneday := 1440; totalMin < 0; totalMin += oneday {
		}
	}
	return Clock{min: totalMin}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60, c.min%60)
}
func (c Clock) Add(min int) Clock {
	totalMin := (c.min + min) % 1440
	if totalMin < 0 {
		for oneday := 1440; totalMin < 0; totalMin += oneday {
		}
	}
	return Clock{min: totalMin}
}
