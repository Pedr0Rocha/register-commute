package storage

import (
	"encoding/json"
	"fmt"
	"sort"

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

func CreateCommutes(newCommutes []c.Commute) error {
	commutes, err := GetCommutes()
	if err != nil {
		return fmt.Errorf("could not create commute, error parsing file: %s", err)
	}

	commutes = append(commutes, newCommutes...)

	sort.Slice(commutes, func(i, j int) bool {
		return commutes[i].Date > commutes[j].Date
	})

	updatedData, err := json.MarshalIndent(commutes, "", " ")
	if err != nil {
		return fmt.Errorf("error registering new commute: %s", err)
	}

	err = WriteToFile(updatedData)
	if err != nil {
		return fmt.Errorf("could not write to file: %s", err)
	}
	return nil
}
