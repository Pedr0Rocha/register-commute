package commute

type Commute struct {
	Id        string `json:"id"`
	Date      string `json:"date"`
	Transport string `json:"transport"`
	CreatedAt string `json:"created_at"`
}

type CommuteMap map[string]Commute
