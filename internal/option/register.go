package option

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/AlecAivazis/survey/v2"
	c "github.com/Pedr0Rocha/register-commute/internal/commute"
	"github.com/Pedr0Rocha/register-commute/internal/storage"
)

const (
	REGISTER_OPTION = "Register new commute"
)

var (
	commutesMap c.CommuteMap
)

func RegisterNewCommute() {
	commutes, err := storage.GetCommutes()
	if err != nil {
		fmt.Println("Not possible to register new commute", err)
		return
	}

	commutesMap = make(c.CommuteMap)
	for _, commute := range commutes {
		commutesMap[commute.Date] = commute
	}

	daysAnswer := askDays()
	transportAnswer := askTransport()

	for _, day := range daysAnswer {
		newEntry := c.Commute{
			Date:      day,
			Transport: transportAnswer,
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

	err = os.WriteFile(storage.FILE_PATH, updatedData, os.ModePerm)
	if err != nil {
		fmt.Println("Could not write into the file:", err)
		return
	}
}

func getTransportationOptions() []string {
	return []string{"Bike", "Public Transport"}
}

func askTransport() string {
	transport := ""
	transportPrompt := &survey.Select{
		Message: "Choose the transportation:",
		Options: getTransportationOptions(),
	}
	survey.AskOne(transportPrompt, &transport, survey.WithValidator(survey.Required))

	return transport
}

func askDays() []string {
	days := []string{}
	var weekdayDateOptions c.CommuteDateSlice = getWeekdayOptions()
	daysPrompt := &survey.MultiSelect{
		Message: "Which days do you want to register?",
		Options: weekdayDateOptions.FormatToOptionTitle(),
		Description: func(value string, index int) string {
			display := weekdayDateOptions[index].FormatToDisplay()
			if _, exists := commutesMap[value]; exists {
				display = display + " [Already registered]"
			}
			return display
		},
	}
	var dateValidator survey.Validator = func(ans interface{}) error {
		pickedDate := ans.([]survey.OptionAnswer)[0].Value
		if _, exists := commutesMap[pickedDate]; exists {
			return fmt.Errorf("date already registered")
		}
		return nil
	}
	validators := survey.ComposeValidators(dateValidator, survey.Required)
	survey.AskOne(daysPrompt, &days, survey.WithValidator(validators))

	return days
}

// getWeekdayOptions calculates all the weekdays according to the current
// date. Week starts on a Sunday.
func getWeekdayOptions() c.CommuteDateSlice {
	currentTime := time.Now()
	startOfWeek := currentTime.AddDate(0, 0, -int(currentTime.Weekday()))

	options := c.CommuteDateSlice{}
	for i := 0; i < 7; i++ {
		options = append(options, c.CommuteDate{Time: startOfWeek})
		startOfWeek = startOfWeek.AddDate(0, 0, 1)
	}

	return options
}
