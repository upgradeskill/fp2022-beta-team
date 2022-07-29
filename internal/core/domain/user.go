package domain

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func NewUser(id uint64, name, role string) *User {
	return &User{
		ID:   id,
		Name: name,
		Role: role,
	}
}
