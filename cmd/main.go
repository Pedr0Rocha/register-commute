package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	c "github.com/Pedr0Rocha/register-commute/internal/commute"
	"github.com/Pedr0Rocha/register-commute/internal/option"
	"github.com/Pedr0Rocha/register-commute/internal/store"
)

func main() {
	// @TODO: move this to each command
	data := store.ReadFromFile()

	var commutes []c.Commute
	err := json.Unmarshal(data, &commutes)
	if err != nil {
		fmt.Println("Could not parse commutes file:", err)
		return
	}

	for {
		mainLoopOption := ""
		mainLoopPrompt := &survey.Select{
			Message: "What do you want to do?",
			Options: []string{option.DISPLAY_OPTION, option.REGISTER_OPTION, "Exit"},
		}
		survey.AskOne(mainLoopPrompt, &mainLoopOption)

		switch mainLoopOption {
		case "Exit":
			os.Exit(0)
		case option.DISPLAY_OPTION:
			option.DisplayCommutes(commutes)
		case option.REGISTER_OPTION:
			option.RegisterNewCommute(commutes)
		}
	}
}
