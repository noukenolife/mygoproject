package user_dto

type Email struct {
	Value string `json:"value"`
}

type UserProfile struct {
	Email Email `json:"email"`
}
