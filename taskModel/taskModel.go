package taskModel

type Task struct {
	ID string `json:"id,omitempty"`
	Todo string `json:"todo,omitempty"`
	User string `json:"user,omitempty"`
}

func (t Task) Add(tasks *[]Task) {
	*tasks = append(*tasks, t)
}

func (task Task) Remove(id string, tasks *[]Task) {
	slice := *tasks

	for index, item := range *tasks {
		if item.ID == id {
			next := index + 1
			slice = append(slice[:index], slice[next:]...)
			*tasks = slice
			break
		}
	}
}
