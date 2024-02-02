package option

import (
	"fmt"

	c "github.com/Pedr0Rocha/register-commute/internal/commute"
	"github.com/Pedr0Rocha/register-commute/internal/storage"
)

const (
	UPDATE_OPTION = "Update a registered commute"
)

func Update() {
	// @TODO: display options to choose and then ask for input again
	commute := c.Commute{
		CreatedAt: "2023-11-21 20:07:32.571157 +0000 UTC",
		Id:        "b7ae95a6-2489-4762-a0ec-a60c255c3c9c",
		Date:      "2023-01-01",
		Transport: "Bike",
	}

	err := storage.UpdateCommute(commute)
	if err != nil {
		fmt.Println("Could not update commute", err)
	}
}
