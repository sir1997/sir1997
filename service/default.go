package service

import (
	"todo/handle"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func EngineStater() *gin.Engine {

	//启动引擎
	r := gin.Default()

	//cookie的使用，[]byte("secret")加密
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))

	//渲染模板
	r.LoadHTMLGlob("./views/*")
    //加载静态模板
	r.Static("./static", "static")

	//登录认证
	r.GET("/", handle.LoginCapthch)

	//注册认证
	r.GET("/logon", handle.LogonSet)

	//导出数据库文章数据
	r.GET("/home", handle.StaterOk)
	
	group := r.Group("/home")
	{
	
	group.GET("/detailed",handle.SetLike)   //文章详情页
	group.GET("/study",handle.ClassifyStudy)                 
	group.GET("/think",handle.ClassifyStudy)                  
	group.GET("/problem",handle.ClassifyStudy)
	group.GET("/bottom",handle.ClassifyStudy)
	group.GET("/model",handle.ClassifyStudy)
	
   }

	//导入文章前端数据到数据库
	r.GET("/text", handle.InsetDate)

	return r

}
