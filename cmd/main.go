package main

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Pedr0Rocha/register-commute/internal/option"
)

func main() {
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
			option.DisplayCommutes()
		case option.REGISTER_OPTION:
			option.RegisterNewCommute()
		}
	}
}
