package model

type News struct {
	Id int          `json:"id" gorm:"column:id"`
	Title string    `json:"title" gorm:"column:title"`
	Content string  `json:"content" gorm:"column:content"`
	CateId int      `json:"cate_id" gorm:"column:cate_id"`
	CreateTime int  `json:"create_time" gorm:"column:create_time"`
}


type NewsAll struct {
	Id int          `json:"id" gorm:"column:id"`
	Title string    `json:"title" gorm:"column:title"`
	Content string  `json:"content" gorm:"column:content"`
	CateId int      `json:"cate_id" gorm:"column:cate_id"`
	CateName string `json:"cate_name" gorm:"column:cate_name"`
}