package todo

import "time"


type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
	UserID string `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

