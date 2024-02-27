package model

type User struct {
	Id       int    `gorm:"type:int;primary_key"`
	UserName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
