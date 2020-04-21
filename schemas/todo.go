package schemas

type Todo struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Date string `json:"date"`
}
