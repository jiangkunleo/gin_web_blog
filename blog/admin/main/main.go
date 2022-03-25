package main

import (
	"github.com/gin-gonic/gin"
	"go_code/blog/admin/router"
)

func main() {
	engine := gin.Default()
	engine.Static("/static", "../static")
	engine.LoadHTMLGlob("../tem/**/*")
	r := router.InitRouter(engine)
	r.Run(":8080")
}
