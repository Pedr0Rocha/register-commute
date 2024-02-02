package tools

import (
	"fmt"

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
	return fmt.Sprintf("Date: %s | Transport: %s %s",
		commute.Date,
		GetTransportEmoji(commute.Transport),
		commute.Transport,
	)
}
