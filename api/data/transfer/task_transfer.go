package transfer

import (
	"noda/api/data/types"
)

/* Transfers a task creation request.  */
type TaskCreation struct {
	Title       string             `json:"title" validate:"required"`
	Headline    string             `json:"headline"`
	Description string             `json:"description"`
	Priority    types.TaskPriority `json:"priority"`
	Status      types.TaskStatus   `json:"status"`
}

func (t *TaskCreation) Validate() error {
	return validate(t)
}

/* Transfers a task update request.  */
type TaskUpdate struct {
	Title       string             `json:"title"`
	Headline    string             `json:"headline"`
	Description string             `json:"description"`
	Priority    types.TaskPriority `json:"priority"`
	Status      types.TaskStatus   `json:"status"`
}
