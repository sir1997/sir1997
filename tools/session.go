package tools

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, key string, value string) error {
	ses := sessions.Default(c)

	if ses.Get(key) == nil {
		ses.Set(key, value)
		return ses.Save()
	}
	return nil
}

func GetSession(c *gin.Context, key string) interface{} {
	ses := sessions.Default(c)
	session := ses.Get(key)
	return session
}
