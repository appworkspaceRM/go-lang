package modeluser

import "time"

type User struct {
	ID          int       `json:"id"`
	FullName    string    `json:"fullName" gorm:"type: varchar(255)"`
	UserName    string    `json:"userName" gorm:"type: varchar(255)"`
	Email       string    `json:"email" gorm:"type: varchar(255)"`
	Password    string    `json:"password" gorm:"type: varchar(255)"`
	Role        string    `json:"-" gorm:"type: varchar(255)"`
	PhoneNumber string    `json:"phoneNumber" gorm:"type: varchar(255)"`
	Gender      string    `json:"gender" gorm:"type: varchar(255)"`
	Address     string    `json:"address" gorm:"type: varchar(255)"`
	CreatedAd   time.Time `json:"-"`
	UpdateAd time.Time `json:"-"`
}

type UserResponse struct {
	ID          int       `json:"id"`
	FullName    string    `json:"fullName" gorm:"type: varchar(255)"`
	UserName    string    `json:"userName" gorm:"type: varchar(255)"`
	Email       string    `json:"email" gorm:"type: varchar(255)"`
	Password    string    `json:"password" gorm:"type: varchar(255)"`
	Role        string    `json:"-" gorm:"type: varchar(255)"`
	PhoneNumber string    `json:"phoneNumber" gorm:"type: varchar(255)"`
	Gender      string    `json:"gender" gorm:"type: varchar(255)"`
	Address     string    `json:"address" gorm:"type: varchar(255)"`
	CreatedAd   time.Time `json:"-"`
	UpdateAd time.Time `json:"-"`
}

func (UserResponse)TableName() string {
	return "users"
}