package domain

type RoleUser struct {
	ID       string `gorm:"type:varchar(100);primary_key"`
	RoleName string `gorm:"type:varchar(100)"`
}
