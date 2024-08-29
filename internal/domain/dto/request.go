package dto

type UserRequest struct {
	Username string `json:"username" gorm:"type:varchar(100)"`
	Email    string `json:"email" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
	Role     string `json:"role" gorm:"type:varchar(100)"`
}
