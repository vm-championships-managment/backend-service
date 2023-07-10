package dtos

type UserDto struct {
	Id        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Birthdate string `json:"birthdate,omitempty"`
	Phone     string `json:"phone,omitempty"`
}
