package domain

type Task struct {
	Id          int    `json:"id,omitempty"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Next        int    `json:"next,omitempty"` //maybe a task later
	Done        bool   `json:"done,omitempty"`
}
