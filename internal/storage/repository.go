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

func DeleteCommute(id string) error {
	commutes, err := GetCommutes()
	if err != nil {
		return err
	}

	indexToRemove := -1
	for i, cm := range commutes {
		if cm.Id == id {
			indexToRemove = i
			break
		}
	}

	if indexToRemove != -1 {
		// commute at indexToRemove gets overwritten by last element
		// commutes is assigned to commutes - last element
		commutes[indexToRemove] = commutes[len(commutes)-1]
		commutes = commutes[:len(commutes)-1]
	}

	err = formatAndWriteToFile(commutes)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCommute(commute c.Commute) error {
	commutes, err := GetCommutes()
	if err != nil {
		return err
	}

	indexToUpdate := -1
	for i, cm := range commutes {
		if cm.Id == commute.Id {
			indexToUpdate = i
			break
		}
	}

	if indexToUpdate != -1 {
		commutes[indexToUpdate] = commute
	}

	err = formatAndWriteToFile(commutes)
	if err != nil {
		return err
	}

	return nil
}

func formatAndWriteToFile(commutes []c.Commute) error {
	sort.Slice(commutes, func(i, j int) bool {
		return commutes[i].Date > commutes[j].Date
	})

	updatedData, err := json.MarshalIndent(commutes, "", " ")
	if err != nil {
		return err
	}

	err = WriteToFile(updatedData)
	if err != nil {
		return err
	}

	return nil
}

func CreateCommutes(newCommutes []c.Commute) error {
	commutes, err := GetCommutes()
	if err != nil {
		return fmt.Errorf("could not create commute, error parsing file: %s", err)
	}

	commutes = append(commutes, newCommutes...)

	err = formatAndWriteToFile(commutes)
	if err != nil {
		return fmt.Errorf("error registering new commute: %s", err)
	}

	return nil
}
