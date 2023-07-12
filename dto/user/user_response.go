package dtouser

type UserResponse struct {
	ID          int    `json:"id"`
	FullName    string `json:"fullName"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	Role        string `json:"-"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
}