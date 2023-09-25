package option

import (
	"fmt"

	c "github.com/Pedr0Rocha/register-commute/internal/commute"
)

const (
	DISPLAY_OPTION = "Check commute days"
)

func DisplayCommutes(commutes []c.Commute) {
	for i, commute := range commutes {
		if i >= 30 {
			break
		}
		fmt.Printf("Date: %s | Transport: %s %s\n",
			commute.Date,
			getTransportEmoji(commute.Transport),
			commute.Transport,
		)
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
