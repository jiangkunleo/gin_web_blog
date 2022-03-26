package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	conf,_ := ParseConfig("../config/admin.json")
	mysqlUser := conf.Database.User
	mysqlPass := conf.Database.Password
	mysqlHost := conf.Database.Host
	mysqlPort := conf.Database.Port
	mysqlDbname := conf.Database.DbName
	mysqlChartset := conf.Database.Chartset
	addr := mysqlUser+":"+mysqlPass+"@("+mysqlHost+":"+mysqlPort+")/"+
		mysqlDbname+"?charset="+mysqlChartset+
		"&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open("mysql", addr)
	if err != nil {
		panic(err)
	}
}
