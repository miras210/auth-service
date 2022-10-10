package models

type User struct {
	ID       int64  `json:"id"`
	PublicID string `json:"public_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpUser struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type SignInUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
