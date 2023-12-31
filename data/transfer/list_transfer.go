package transfer

/* Transfers a list creation request.  */
type ListCreation struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (l *ListCreation) Validate() error {
	return validate(l)
}

/* Transfers a list update request.  */
type ListUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
