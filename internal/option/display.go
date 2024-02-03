package option

import (
	"fmt"

	c "github.com/Pedr0Rocha/register-commute/internal/commute"
	"github.com/Pedr0Rocha/register-commute/internal/storage"
	"github.com/Pedr0Rocha/register-commute/tools"
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

	periodCount := make(map[string]int)
	for _, commute := range commutes {
		periodCount[buildPeriodKey(commute)]++
	}

	monthsDisplayed := 0
	currMonth := ""
	for _, commute := range commutes {
		periodKey := buildPeriodKey(commute)

		if currMonth != commute.GetMonth() {
			monthsDisplayed++
			if monthsDisplayed > DISPLAY_MONTH_COUNT {
				break
			}

			fmt.Printf("====> %s/%s [%d] <====\n",
				commute.GetMonth(),
				commute.GetYear(),
				periodCount[periodKey],
			)
		}

		currMonth = commute.GetMonth()
		fmt.Printf(tools.GetDefaultCommuteDisplay(commute) + "\n")
	}
}

func buildPeriodKey(commute c.Commute) string {
	year := commute.GetYear()
	month := commute.GetMonth()
	return fmt.Sprintf("%s%s", year, month)
}
