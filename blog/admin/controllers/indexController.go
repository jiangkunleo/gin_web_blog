package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go_code/blog/admin/model"
	"go_code/blog/admin/utils"
	"net/http"
)

type Index struct {}

func (i *Index) OutLogin(c *gin.Context) {
	store := sessions.NewCookieStore([]byte("secrect"))
	session,_ := store.Get(c.Request,"user")
	session.Values["is_login"] = false
	session.Save(c.Request,c.Writer)
	c.JSON(http.StatusOK, gin.H{"code": 1})
}

func (i *Index) ShowIndex(c *gin.Context) {
	store := sessions.NewCookieStore([]byte("secrect"))
	session,_ := store.Get(c.Request,"user")
	username := session.Values["user_name"]
	c.HTML(http.StatusOK, "admin/index.html", gin.H{"username":username})
}
func (i *Index) CateIndex(c *gin.Context) {
	store := sessions.NewCookieStore([]byte("secrect"))
	session,_ := store.Get(c.Request,"user")
	username := session.Values["user_name"]

	//获取分类列表数据并树状排序
	var cates []model.Cate
	Db := utils.Db
	Db.Table("cate").Find(&cates)
	var menus []utils.Menu
	for _,v := range cates {
		var menu utils.Menu
		menu.Id = v.Id
		menu.Pid = v.Pid
		menu.CateName = v.CateName
		menu.Desc = v.Desc
		menu.CreateTime = v.CreateTime
		menus = append(menus,menu)
	}
	newMenus := utils.GetMenu(menus,0)

	c.HTML(http.StatusOK, "admin/cate_index.html", gin.H{
		"username":username,
		"list":newMenus,
	})
}

func (i *Index) ShowAddCate(c *gin.Context) {
	store := sessions.NewCookieStore([]byte("secrect"))
	session,_ := store.Get(c.Request,"user")
	username := session.Values["user_name"]
	c.HTML(http.StatusOK, "admin/cate_add.html", gin.H{"username":username})
}

