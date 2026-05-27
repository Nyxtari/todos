package todos

import "time"

type Todo struct {
	Id          int32     `json:"id"`
	IsDone      bool      `json:"isDone"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Difficulty  int32     `json:"difficulty"`
}
