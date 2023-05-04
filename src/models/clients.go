package models

type Clientstruct struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	JobTitle string `json:"jobTitle,omitempty"`
	Role     string `json:"role,omitempty"`
}
