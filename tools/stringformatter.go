package tools

import (
	"fmt"
	"time"

	c "github.com/Pedr0Rocha/register-commute/internal/commute"
)

func GetTransportEmoji(transport string) string {
	switch transport {
	case "Bike":
		return "🚲"
	case "Public Transport":
		return "🚈"
	default:
		return ""
	}
}

func GetDefaultCommuteDisplay(commute c.Commute) string {
	date, err := time.Parse("2006-01-02", commute.Date)
	if err != nil {
		fmt.Println("Dates are not formatted propertly in the file:", err)
		return ""
	}

	// Pad empty spaces to align the text better in the terminal
	maxLenght := 10
	offset := maxLenght - len(date.Weekday().String())
	offsetString := ""
	for i := 0; i < offset; i++ {
		offsetString += " "
	}

	return fmt.Sprintf("Date: %s (%s)%s| Transport: %s %s",
		commute.Date,
		date.Weekday(),
		offsetString,
		GetTransportEmoji(commute.Transport),
		commute.Transport,
	)
}
