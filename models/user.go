package models

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Username string `gorm:"type:varchar(300)" json:"username"`
	Password string `gorm:"type:varchar(300)" json:"password"`
	Status   int    `gorm:"type:integer" json:"status	"`
}
