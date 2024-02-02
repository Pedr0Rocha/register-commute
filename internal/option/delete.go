package option

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	c "github.com/Pedr0Rocha/register-commute/internal/commute"
	"github.com/Pedr0Rocha/register-commute/internal/storage"
	"github.com/Pedr0Rocha/register-commute/tools"
)

const (
	DELETE_OPTION          = "Delete a registered commute"
	LIMIT_PAST_DAYS_DELETE = 30
)

func Delete() {
	commutes, err := storage.GetCommutes()
	if err != nil {
		fmt.Println("Could not fetch the current commutes", err)
		return
	}

	commutesMap = make(c.CommuteMap)
	commutesIds := []string{}
	for i, commute := range commutes {
		if i >= LIMIT_PAST_DAYS_DELETE {
			break
		}

		commutesMap[commute.Id] = commute
		commutesIds = append(commutesIds, commute.Id)
	}

	idsToDelete := []string{}
	deletePrompt := &survey.MultiSelect{
		Message: "Which day you want to delete?",
		Options: commutesIds,
		VimMode: true,
		Description: func(value string, index int) string {
			return tools.GetDefaultCommuteDisplay(commutesMap[value])
		},
	}
	survey.AskOne(deletePrompt, &idsToDelete)

	for _, id := range idsToDelete {
		err = storage.DeleteCommute(id)
		if err != nil {
			fmt.Println("Could not delete commute", err)
		}
	}
}
