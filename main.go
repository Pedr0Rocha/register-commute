package main

import (
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

// buildDayOptions calculates all the weekdays according to the current
// date and returns a array of strings containing the whole week formatted
// as: Saturday 22/09.
// Week starts on a Sunday.
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
	daysPrompt := &survey.MultiSelect{
		Message: "Which days do you want to register?",
		Options: buildDayOptions(),
	}
	survey.AskOne(daysPrompt, &days, survey.WithValidator(survey.Required))

	var transport string
	transportPrompt := &survey.Select{
		Message: "Choose the transportation:",
		Options: []string{"Bike", "Public Transport"},
	}
	survey.AskOne(transportPrompt, &transport, survey.WithValidator(survey.Required))

	fmt.Printf("%s - %s\n", days, transport)
}
