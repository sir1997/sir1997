package model

import "time"

//评论的实现
type Comment struct {
	Id int `json:"id"`
	//根据会话时间查询相应数据
	Time     time.Time `json:"time"`
	Username string    `json:"username"`
	//显示会话内容
	Comment string `json:"comment"`
}
