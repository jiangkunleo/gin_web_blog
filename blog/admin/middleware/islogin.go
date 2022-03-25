package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context){
		store := sessions.NewCookieStore([]byte("secrect"))
		session,err := store.Get(c.Request,"user")
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect ,"/login/showLogin")
			return
		}
		isLogin := session.Values["is_login"]

		if isLogin == false {
			c.Redirect(http.StatusTemporaryRedirect ,"/login/showLogin")
			return
		}

	}
}
