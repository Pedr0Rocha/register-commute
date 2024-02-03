package main

import (
	"os"

	"github.com/Pedr0Rocha/register-commute/internal/option"
)

func main() {
	for {
		mainLoopOption := option.MainLoop()

		switch mainLoopOption {
		case "Exit":
			os.Exit(0)
		case option.DISPLAY_OPTION:
			option.DisplayCommutes()
		case option.REGISTER_OPTION:
			option.RegisterNewCommute()
		case option.DELETE_OPTION:
			option.Delete()
		case option.STATISTICS_OPTION:
			option.ShowStats()
		}
	}
}
