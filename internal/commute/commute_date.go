package commute

import (
	"time"
)

type CommuteDate struct {
	time.Time
}

type CommuteDateSlice []CommuteDate

func (d CommuteDate) FormatToDisplay() string {
	return d.Weekday().String() + " " + d.Format("02/01")
}

func (d CommuteDate) formatToRegister() string {
	return d.Format("2006-01-02")
}

func (days CommuteDateSlice) FormatToOptionTitle() []string {
	displayOptions := []string{}
	for i := 0; i < len(days); i++ {
		displayOptions = append(displayOptions, days[i].formatToRegister())
	}
	return displayOptions
}
