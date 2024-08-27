package domain

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"type:varchar(100);primaryKey"`
	Username  string    `json:"username" gorm:"type:varchar(100)"`
	Email     string    `json:"email" gorm:"type:varchar(100)"`
	Password  string    `json:"password" gorm:"type:varchar(100)"`
	Role      string    `json:"role" gorm:"type:varchar(50)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserDetail struct {
	Id          string `json:"id"`
	FullName    string `json:"full_name"`
	Age         int    `json:"age"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
