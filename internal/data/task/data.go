package task

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CompletedAt string `json:"completed_at"`
}
