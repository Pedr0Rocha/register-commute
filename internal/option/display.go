package option

import (
	"fmt"
	"time"

	c "github.com/Pedr0Rocha/register-commute/internal/commute"
	"github.com/Pedr0Rocha/register-commute/internal/storage"
)

const (
	DISPLAY_OPTION      = "Check commute days"
	DISPLAY_MONTH_COUNT = 3
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

	month := make([]string, 0, len(displayMap))
	for k := range displayMap {
		month = append(month, k)
	}

	for i, month := range month {
		if i >= DISPLAY_MONTH_COUNT {
			break
		}
		fmt.Printf("====> %s <====\n", month)
		for _, commute := range displayMap[month] {
			fmt.Printf("Date: %s | Transport: %s %s\n",
				commute.Date,
				getTransportEmoji(commute.Transport),
				commute.Transport,
			)
		}
	}
}

func getTransportEmoji(transport string) string {
	switch transport {
	case "Bike":
		return "🚲"
	case "Public Transport":
		return "🚈"
	default:
		return ""
	}
}
