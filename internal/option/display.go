package option

import (
	"fmt"

	"github.com/Pedr0Rocha/register-commute/internal/storage"
)

const (
	DISPLAY_OPTION = "Check commute days"
	DISPLAY_COUNT  = 30
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

	for i, commute := range commutes {
		if i >= DISPLAY_COUNT {
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
