package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

type CommuteEntry struct {
	Date      string `json:"date"`
	Transport string `json:"transport"`
}

type CommuteDate struct {
	time.Time
}

type CommuteDateSlice []CommuteDate

func (d CommuteDate) formatToDisplay() string {
	return d.Weekday().String() + " " + d.Format("02/01")
}

func (d CommuteDate) formatToRegister() string {
	return d.Format("2006-01-02")
}

func (days CommuteDateSlice) getDisplayTitles() []string {
	displayOptions := []string{}
	for i := 0; i < len(days); i++ {
		displayOptions = append(displayOptions, days[i].formatToRegister())
	}
	return displayOptions
}

func getTransportEmoji(transport string) string {
	switch transport {
	case "Bike":
		return "🚲"
	case "Public Transport":
		return "🚈"
	default:
		return ""
	}
}

func displayCommutes(commutes []CommuteEntry) {
	for _, commute := range commutes {
		fmt.Printf("Date: %s | Transport: %s %s\n",
			commute.Date,
			getTransportEmoji(commute.Transport),
			commute.Transport,
		)
	}
}

// getWeekdayOptions calculates all the weekdays according to the current
// date. Week starts on a Sunday.
func getWeekdayOptions() []CommuteDate {
	currentTime := time.Now()
	startOfWeek := currentTime.AddDate(0, 0, -int(currentTime.Weekday()))

	options := []CommuteDate{}
	for i := 0; i < 7; i++ {
		options = append(options, CommuteDate{startOfWeek})
		startOfWeek = startOfWeek.AddDate(0, 0, 1)
	}

	return options
}

func getTransportationOptions() []string {
	return []string{"Bike", "Public Transport"}
}

func main() {
	filePath := "commutes.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Could not find file.")
		return
	}

	var commutes []CommuteEntry
	err = json.Unmarshal(data, &commutes)
	if err != nil {
		fmt.Println("Could not parse commutes file:", err)
		return
	}

	for {
		mainLoopOption := ""
		mainLoopPrompt := &survey.Select{
			Message: "What do you want to do?",
			Options: []string{"Check commute days", "Register new commute", "Exit"},
		}
		survey.AskOne(mainLoopPrompt, &mainLoopOption)

		switch mainLoopOption {
		case "Exit":
			os.Exit(0)
		case "Check commute days":
			displayCommutes(commutes)
		case "Register new commute":
			days := []string{}
			var weekdayDateOptions CommuteDateSlice = getWeekdayOptions()
			daysPrompt := &survey.MultiSelect{
				Message: "Which days do you want to register?",
				Options: weekdayDateOptions.getDisplayTitles(),
				Description: func(value string, index int) string {
					return weekdayDateOptions[index].formatToDisplay()
				},
			}
			survey.AskOne(daysPrompt, &days, survey.WithValidator(survey.Required))

			transport := ""
			transportPrompt := &survey.Select{
				Message: "Choose the transportation:",
				Options: getTransportationOptions(),
			}
			survey.AskOne(transportPrompt, &transport, survey.WithValidator(survey.Required))

			fmt.Printf("%s - %s\n", days, transport)

			for _, day := range days {
				newEntry := CommuteEntry{
					Date:      day,
					Transport: transport,
				}

				commutes = append(commutes, newEntry)
			}
			sort.Slice(commutes, func(i, j int) bool {
				return commutes[i].Date > commutes[j].Date
			})

			updatedData, err := json.MarshalIndent(commutes, "", " ")
			if err != nil {
				fmt.Println("Error registering new commute:", err)
				return
			}

			err = os.WriteFile(filePath, updatedData, os.ModePerm)
			if err != nil {
				fmt.Println("Could not write into the file:", err)
				return
			}
		}
	}
}
