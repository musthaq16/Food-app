package model

type OneTimePassword struct {
	Id     int    `gorm:"type:int;primary_key" json:"id"`
	Otp    string `gorm:"uniqueIndex;type:string" json:"otp"`
	UserId int    `gorm:"foreignkey:id" json:"user_id"`
}
