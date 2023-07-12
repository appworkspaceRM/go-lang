package dtouser

type UserResponse struct {
	ID          int    `json:"id"`
	FullName    string `json:"fullName" gorm:"type: varchar(255)"`
	UserName    string `json:"userName" gorm:"type: varchar(255)"`
	Email       string `json:"email" gorm:"type: varchar(255)"`
	Password    string `json:"-" gorm:"type: varchar(255)"`
	Role        string `json:"-" gorm:"type: varchar(255)"`
	PhoneNumber string `json:"phoneNumber" gorm:"type: varchar(255)"`
	Gender      string `json:"gender" gorm:"type: varchar(255)"`
	Address     string `json:"address" gorm:"type: varchar(255)"`
}