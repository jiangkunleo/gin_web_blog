package controllers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"go_code/blog/admin/model"
	"go_code/blog/admin/utils"
	"net/http"
	"time"
)

type Api struct {}



func (a *Api) Test(c *gin.Context) {
	var news []model.News
	db := utils.Db
	db.Table("news").Find(&news)

	//连表查询
	//var news []model.NewsAll
	//db := utils.Db
	//db.Table("news n").
	//	Select("n.id,n.title,n.content,n.cate_id,c.cate_name").
	//	Joins("left join cate c on n.cate_id = c.id ").Scan(&news)
	//c.JSON(http.StatusOK, gin.H{"list":news})
	//path := "../config/admin.json"

	//读取配置文件内容
	//path2, _ := filepath.Abs("../config/admin.json")
	//path := "../config/admin.json"
	//config,_ := utils.ParseConfig(path)
	//c.JSON(http.StatusOK, gin.H{"config":config,"path":path2})

	//redis 连接池操作
	redisPool := utils.Pool //获取连接池（main值已经初始化好）
	rdb := redisPool.Get()  //连接池中取出一个连接
	rdb.Do("Set","abc",200) //redis 设置值
	r,_ := redis.Int(rdb.Do("Get","abc")) //redis 取出值
	fmt.Println(r)
	defer redisPool.Close() //关闭连接池

	//获取当前时间戳
	chuo := 1648271731
	tm := time.Unix(int64(chuo),0)
	timeStr := tm.Format("2006-01-02 15:04:05")

	tim ,_ :=  time.Parse("2006-01-02 15:04:05","2022-02-01 10:45:15")
	chuo2 := tim.Unix()
	c.JSON(http.StatusOK, gin.H{"时间":timeStr,"时间戳":chuo2,"reids abc":r,"news":news})
}