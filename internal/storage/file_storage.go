package storage

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	FILE_PATH = "commutes.json"
)

func ReadFromFile() []byte {
	data, err := os.ReadFile(FILE_PATH)

	if errors.Is(err, os.ErrNotExist) {
		initFileStorage()
	} else if err != nil {
		log.Fatal(err)
	}
	return data
}

func WriteToFile(bytes []byte) error {
	err := os.WriteFile(FILE_PATH, bytes, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not write into the file: %s", err)
	}
	return nil
}

func initFileStorage() {
	fmt.Println("File not found. Creating new file to register commutes")
	f, err := os.Create(FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
