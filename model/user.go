package model

type User struct {
	Id       int    `gorm:"type:int;primary_key" json:"id"`
	UserName string `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
