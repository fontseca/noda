package transfer

/* Transfers a group creation request.  */
type GroupCreation struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (g *GroupCreation) Validate() error {
	return validate(g)
}

/* Transfers a group update request.  */
type GroupUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
