package models

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Fullname   string `json:"full_name"`
	Email      string `json:"email"`
	IsVerified string `json:"is_verified"`
	CreatedAt  string `json:"created_at"`
}
