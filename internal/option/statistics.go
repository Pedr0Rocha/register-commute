package option

import (
	"fmt"

	"github.com/Pedr0Rocha/register-commute/internal/storage"
	"github.com/Pedr0Rocha/register-commute/tools"
)

const (
	STATISTICS_OPTION = "Show statistics"
)

var stats Stats

type Stats struct {
	TotalPerYear      map[string]int
	TotalPerMonth     map[string]int
	TotalPerWeekday   map[string]int
	TotalPerTransport map[string]int
}

func ShowStats() {
	commutes, err := storage.GetCommutes()
	if err != nil {
		fmt.Println("Not possible to display commutes:", err)
		return
	}

	yearCount := make(map[string]int)
	monthCount := make(map[string]int)
	weekdayCount := make(map[string]int)
	transportCount := make(map[string]int)
	for _, commute := range commutes {
		yearCount[commute.GetYear()]++
		monthCount[commute.GetMonth()]++
		weekdayCount[commute.GetWeekday()]++
		transportCount[commute.Transport]++
	}

	stats := Stats{
		TotalPerYear:      yearCount,
		TotalPerMonth:     monthCount,
		TotalPerWeekday:   weekdayCount,
		TotalPerTransport: transportCount,
	}

	stats.display()
}

func (s Stats) display() {
	fmt.Println("====> Year Count <====")
	for k, v := range s.TotalPerYear {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Println("====>  Month Count <====")
	for k, v := range s.TotalPerMonth {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Println("====> Weekday Count <====")
	for k, v := range s.TotalPerWeekday {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Println("====> Transport Count <====")
	for k, v := range s.TotalPerTransport {
		fmt.Printf("%s: %s %d\n", k, tools.GetTransportEmoji(k), v)
	}
}
