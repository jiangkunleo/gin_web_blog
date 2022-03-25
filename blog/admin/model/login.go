package model

type User struct {
	UserName string `json:"username" form:"username" binding:"required,gte=4,lte=10" gorm:"column:username"`
	PassWord string `json:"password" form:"password" binding:"required,gte=6" gorm:"column:password"`
	Passcode string `json:"passcode" form:"passcode" binding:"required"`
}
