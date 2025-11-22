package tasks

import "time"

type Task struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Command   string    `json:"command"`
	Schedule  string    `json:"schedule"`
	CreatedAt time.Time `json:"created_at"`
}
