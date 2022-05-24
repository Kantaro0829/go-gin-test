package json

type CreateTaskJson struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is-done"`
}

type UpdateTaskJson struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is-done"`
}
