package model

type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
