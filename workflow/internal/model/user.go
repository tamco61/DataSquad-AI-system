package model

// type User struct {
// 	ID    string `json:"id"`
// 	Email string `json:"email"`
// 	Name  string `json:"name"`
// }

type UserDTO struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Departament string `json:"departament"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
