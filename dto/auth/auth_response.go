package dtoauth

type AuthRequest struct {
	FullName    string `json:"fullName" form:"fullname" validate:"required"`
	UserName    string `json:"username" form:"username" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
	Role        string `json:"-" form:"role" validate:"required"`
	PhoneNumber string `json:"phoneNumber" form:"phone_number" validate:"required"`
	Gender      string `json:"gender" form:"gender" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
}

type AuthLogin struct {
	UserName string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginResponse struct {
	FullName    string `json:"fullName"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"-"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Token       string `json:"token"`
}