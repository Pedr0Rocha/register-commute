package storage

import (
	"encoding/json"
	"fmt"

	c "github.com/Pedr0Rocha/register-commute/internal/commute"
)

func GetCommutes() ([]c.Commute, error) {
	data := ReadFromFile()

	var commutes []c.Commute
	err := json.Unmarshal(data, &commutes)
	if err != nil {
		fmt.Println("Could not parse commutes file:", err)
		return nil, err
	}

	return commutes, nil
}
