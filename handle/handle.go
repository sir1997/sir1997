package handle

import (
	"fmt"
	"net/http"
	"todo/dao"
	"todo/tools"
	"github.com/gin-gonic/gin"
)

var (
	CookiesDebug string //解决cookie的问题
	PageNext     = 4    //数据库文章分页查询,只展示4页，初始查询为4条数据，查询自增加4
)

//登录
func LoginCapthch(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)

	name := c.Query("username")

	password := c.Query("password")

	//完成sessions操作
	tools.SetSession(c, "key", name)

	str := tools.GetSession(c, "key")

	strs := fmt.Sprintf("%v", str)

	CookiesDebug = strs

	if name != "" && password != "" {

		say := dao.SetLogin(name)
		for _, db := range say {
			if db.Id > 0 {
				if db.Username == name && db.Password == password {
					c.JSON(200, nil)
				} else if db.Password != password {
					c.JSON(200, gin.H{
						"falet": "密码不正确",
					})
				}
			} else {
				c.JSON(200, gin.H{
					"login": "账号未注册,请注册账号",
				})
			}
		}
	}
}

//注册
func LogonSet(c *gin.Context) {
	c.HTML(200, "logon.html", nil)

	name := c.Query("username")
	password := c.Query("password")
	repassword := c.Query("repassword")

	names := dao.LogonRepise(name)

	if names == nil && name != "" && password != "" {

		if password == repassword {
			dao.LogonUser(name, password)
		} else {
			fmt.Println(name)
			fmt.Println(password)
		}

	}

}

//插入文章数据到数据库
func InsetDate(c *gin.Context) {
	//获取前端输入的文章保存到数据库
	c.HTML(http.StatusOK, "text.html", nil)
	title := c.Query("title")
	content := c.Query("content")
	classify := c.Query("sex")

	author := CookiesDebug

	if title != "" && content != "" && classify != "" {

		useid := fmt.Sprintf("%v", author)

		dao.SetArticle(title, useid, content, classify)
	}

   authors :=  c.Query("author")
  communcation := c.Query("communcation")

  com,user := dao.Communcation(communcation,authors)
  //当不是接收对象时，无法查询，查询对应cookie的会话对象
  if CookiesDebug==user{
	c.JSON(200,gin.H{
		"com":com,
		"author":user,
	})
  }

  dao.CollectShow(CookiesDebug)
}

//文章详情页
func SetLike(c *gin.Context) {

	title := c.Query("title")
	username :=c.Query("username")

	like := c.Query("like")
	
	comment := c.Query("comment")

	if like==""&&comment==""{
		
		err := dao.ViewCount(title)
	if err !=nil{
		fmt.Println("自增浏览失败")
	}
	article := dao.OneArticle(title)
	coun := dao.CommnetCount()
	comm:=dao.Comment(CookiesDebug,comment)
		for _, val := range article {
			c.HTML(200,"detailed.html",gin.H{
				 "val":val,
				 "comment":comm,
				"commentcount":coun,
			})
		}

	} else if like!=""&&comment==""{
		errs := dao.LikeCount(title,like)
		
		article := dao.OneArticle(title)
		for _, val := range article {
            coun := dao.CommnetCount()
			comm:=dao.Comment(CookiesDebug,comment)
			c.HTML(200, "detailed.html", gin.H{
				"val":val,
				"comment":comm,
				"commentcount":coun,
			})
		}
	        if errs !=nil{
		         fmt.Println("点赞失败")
	          }
	}
	if comment!=""{
			article := dao.OneArticle(title)
	
			for _, val := range article {
				
				name := "小哈"
				comm:=dao.Comment(name,comment)
				coun := dao.CommnetCount()

				c.HTML(200,"detailed.html",gin.H{
					 "val":val,
					 "comment":comm,
					"commentcount":coun,
				})
			}
		}

 dao.Collect(title,username)
}

//加载文章到前端,并且分页
func StaterOk(c *gin.Context) {

	nextpage := c.Query("pagenext")
	priverpage := c.Query("priverpage")
	
	Arti := dao.Outarticle()
	
	if nextpage == "" && priverpage == "" {

		c.HTML(http.StatusOK, "home.html", Arti[0:PageNext])
	} else if nextpage != "" {

		PageNext = PageNext+4
		if PageNext >= len(Arti) {
			PageNext = 0
		}

		article := Arti[PageNext:PageNext+4]

		c.HTML(http.StatusOK, "home.html", article)
	}

	if priverpage != "" {
		PageNext=PageNext-4
		if PageNext < 0 {
			PageNext = len(Arti) - 1
		}
		c.HTML(http.StatusOK, "home.html", Arti[PageNext:PageNext+4])
	}

}


//分类文章查找
func ClassifyStudy(c *gin.Context) {

	classify := c.Query("classify")

	//文章查找
	classifyarticle := dao.ClassifyFinde(classify)
	
		c.HTML(200, "go_problem.html", classifyarticle)
	

}
