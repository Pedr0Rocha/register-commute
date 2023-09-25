package option

import "github.com/AlecAivazis/survey/v2"

func MainLoop() string {
	mainLoopOption := ""
	mainLoopPrompt := &survey.Select{
		Message: "What do you want to do?",
		Options: []string{DISPLAY_OPTION, REGISTER_OPTION, "Exit"},
	}
	survey.AskOne(mainLoopPrompt, &mainLoopOption)

	return mainLoopOption
}
