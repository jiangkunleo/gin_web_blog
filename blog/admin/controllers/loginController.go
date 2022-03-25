package controllers

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go_code/blog/admin/model"
	"go_code/blog/admin/utils"
	"go_code/blog/admin/validator"
	"net/http"
	"time"
)

type Login struct {}

func (l *Login) Captcha(c *gin.Context) {
	width,height := 80,32
	captchaId := captcha.NewLen(4)
	store := sessions.NewCookieStore([]byte("secrect"))
	session,_ := store.New(c.Request,"captcha")
	session.Values["captcha_id"] = captchaId
	session.Save(c.Request,c.Writer)

	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")

	var content bytes.Buffer
	c.Writer.Header().Set("Content-Type", "image/png")
	_ = captcha.WriteImage(&content, captchaId, width, height)
	http.ServeContent(c.Writer,c.Request,captchaId+".png",time.Time{},bytes.NewReader(content.Bytes()))
}


func (l *Login) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK,"login/login.html",gin.H{})
}

func (l *Login) ToLogin(c *gin.Context) {
	validator.Init()
	var user model.User
	//接收參數
	if err := c.ShouldBind(&user); err != nil {
		errs := validator.Translate(err)
		s := ""
		for _,v := range errs {
			s += v[0]
		}
		c.JSON(http.StatusOK, gin.H{"code":0,"error": s})
		return
	}
	cstore := sessions.NewCookieStore([]byte("secrect"))
	csession,_ := cstore.New(c.Request,"captcha")
	captchaId := csession.Values["captcha_id"]

	if !captcha.VerifyString(captchaId.(string),user.Passcode)  {
		c.JSON(http.StatusOK, gin.H{"code":0,"error": "验证码错误"})
		return
	}

	var dbUser model.User
	utils.Db.Table("user").Where("username = ?",user.UserName).First(&dbUser)
	if dbUser.PassWord != user.PassWord {
		c.JSON(http.StatusOK, gin.H{"code":0,"error": "密码错误"})
		return
	} else {
		store := sessions.NewCookieStore([]byte("secrect"))
		session,_ := store.New(c.Request,"user")
		session.Values["is_login"] = true
		session.Values["user_name"] = user.UserName
		session.Save(c.Request,c.Writer)
	}
	c.JSON(http.StatusOK, gin.H{"code": 1})
}
