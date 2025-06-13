package models

type User struct {
	UserID   uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
