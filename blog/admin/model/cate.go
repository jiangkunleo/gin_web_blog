package model

type Cate struct {
	Id int          `json:"id" form:"id"  gorm:"column:id"`
	Pid int         `json:"pid" form:"pid"  gorm:"column:pid"`
	CateName string `json:"cate_name" form:"cate_name"  gorm:"column:cate_name"`
	Desc string     `json:"desc" form:"desc" gorm:"column:desc"`
	CreateTime int  `json:"create_time" form:"create_time" gorm:"column:create_time"`
}
