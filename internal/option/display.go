package option

import (
	"fmt"

	"github.com/Pedr0Rocha/register-commute/internal/storage"
)

const (
	DISPLAY_OPTION = "Check commute days"
)

func DisplayCommutes() {
	commutes, err := storage.GetCommutes()
	if err != nil {
		fmt.Println("Not possible to display commutes", err)
		return
	}

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
