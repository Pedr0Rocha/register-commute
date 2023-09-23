package main

import (
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

var transportationQuestion = []*survey.Question{
	{
		Name: "transportation",
		Prompt: &survey.Select{
			Message: "Choose the transportation:",
			Options: []string{"Bike", "Public Transport"},
		},
		Validate: survey.Required,
	},
}

func buildDayOptions() []string {
	currentTime := time.Now()
	startOfWeek := currentTime.AddDate(0, 0, -int(currentTime.Weekday()))

	options := []string{}
	for i := 0; i < 7; i++ {
		stringOption := startOfWeek.Weekday().String() + " " + startOfWeek.Format("02/01")
		options = append(options, stringOption)
		startOfWeek = startOfWeek.AddDate(0, 0, 1)
	}

	return options
}

func main() {
	days := []string{}
	prompt := &survey.MultiSelect{
		Message: "Which days do you want to register?",
		Options: buildDayOptions(),
	}
	survey.AskOne(prompt, &days)

	answers := struct {
		Transportation string
	}{}
	err := survey.Ask(transportationQuestion, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s - %s\n", days, answers.Transportation)
}
