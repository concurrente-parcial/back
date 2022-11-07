package models

type User struct {
	// Id          primitive.ObjectID `json:"id,omitempty"`
	FirstName   string `json:"name,omitempty" validate:"required"`
	LastName    string `json:"lastname,omitempty" validate:"required"`
	Email       string `json:"email,omitempty" validate:"required"`
	Password    string `json:"pass,omitempty" validate:"required"`
	Permissions []int  `json:"permissions,omitempty" validate:"required"`
}
