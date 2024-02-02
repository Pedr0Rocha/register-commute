package option

import (
	"fmt"
	"sort"
	"time"

	c "github.com/Pedr0Rocha/register-commute/internal/commute"
	"github.com/Pedr0Rocha/register-commute/internal/storage"
	"github.com/Pedr0Rocha/register-commute/tools"
)

const (
	DISPLAY_OPTION      = "Check commute days"
	DISPLAY_MONTH_COUNT = 12
)

func DisplayCommutes() {
	commutes, err := storage.GetCommutes()
	if err != nil {
		fmt.Println("Not possible to display commutes:", err)
		return
	}

	if len(commutes) == 0 {
		fmt.Println("No commutes yet. Register one to begin.")
		return
	}

	displayMap := make(map[string][]c.Commute)
	for _, commute := range commutes {
		date, err := time.Parse("2006-01-02", commute.Date)
		if err != nil {
			fmt.Println("Dates are not formatted propertly in the file:", err)
			return
		}

		month := date.Month().String()
		displayMap[month] = append(displayMap[month], commute)
	}

	months := make([]string, 0, len(displayMap))
	for k := range displayMap {
		months = append(months, k)
	}

	sort.Slice(months, func(i, j int) bool {
		return monthMap[months[i]] > monthMap[months[j]]
	})

	for i, month := range months {
		if i >= DISPLAY_MONTH_COUNT {
			break
		}

		fmt.Printf("====> %s [%d] <====\n", month, len(displayMap[month]))

		for _, commute := range displayMap[month] {
			fmt.Printf(tools.GetDefaultCommuteDisplay(commute) + "\n")
		}
	}
}

var monthMap = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}
