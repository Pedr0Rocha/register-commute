package commute

import (
	"strconv"
	"time"
)

type Commute struct {
	Id        string `json:"id"`
	Date      string `json:"date"`
	Transport string `json:"transport"`
	CreatedAt string `json:"created_at"`
}

type CommuteMap map[string]Commute

func (c Commute) GetDateAsTime() time.Time {
	date, _ := time.Parse("2006-01-02", c.Date)
	return date
}

func (c Commute) GetYear() string {
	date := c.GetDateAsTime()
	return strconv.Itoa(date.Year())
}

func (c Commute) GetMonth() string {
	date := c.GetDateAsTime()
	return date.Month().String()
}

func (c Commute) GetWeekday() string {
	date := c.GetDateAsTime()
	return date.Weekday().String()
}
