package user

type User struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
