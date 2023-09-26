package option

import (
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
	c "github.com/Pedr0Rocha/register-commute/internal/commute"
	"github.com/Pedr0Rocha/register-commute/internal/storage"
	"github.com/google/uuid"
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
	if len(daysAnswer) <= 0 {
		return
	}

	transportAnswer := askTransport()

	newCommutes := []c.Commute{}
	for _, day := range daysAnswer {
		newCommutes = append(newCommutes, c.Commute{
			Id:        uuid.New().String(),
			Date:      day,
			Transport: transportAnswer,
			CreatedAt: time.Now().UTC().String(),
		})
	}

	err = storage.CreateCommutes(newCommutes)
	if err != nil {
		fmt.Println(err)
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
		optionAnswer := ans.([]survey.OptionAnswer)
		if len(optionAnswer) <= 0 {
			return nil
		}

		pickedDate := optionAnswer[0].Value
		if _, exists := commutesMap[pickedDate]; exists {
			return fmt.Errorf("date already registered")
		}
		return nil
	}
	validators := survey.ComposeValidators(dateValidator)
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
