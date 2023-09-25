package commute

type Commute struct {
	Date      string `json:"date"`
	Transport string `json:"transport"`
}

type CommuteMap map[string]Commute
