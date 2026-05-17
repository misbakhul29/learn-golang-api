package dto

type InputUser struct {
	Name  *string `json:"name,omitempty" doc:"Name of the user"`
	Email *string `json:"email,omitempty" doc:"Email of the user"`
}
