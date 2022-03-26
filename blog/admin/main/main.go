package main

import (
	"github.com/gin-gonic/gin"
	"go_code/blog/admin/router"
	"go_code/blog/admin/utils"
)

func main() {
	utils.InitRedisPool() //初始化连接池
	engine := gin.Default()
	engine.Static("/static", "../static")
	engine.LoadHTMLGlob("../tem/**/*")
	r := router.InitRouter(engine)
	r.Run(":8080")
}
