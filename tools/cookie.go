package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CookieSetService(c *gin.Context) {
	cookieName := &http.Cookie{
		Name:"sessionid",
            Value:"1jl3dndo0ex82kal",
            MaxAge:60*60,
            //Expires: expires,
            Domain:"127.0.0.1:8084",
            Path:"/",
            HttpOnly:true,
	}
	http.SetCookie(c.Writer, cookieName)
}
