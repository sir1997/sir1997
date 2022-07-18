package model

import "time"

type LoginStruct struct {
	Id           int       `json:"id" form:"id"`
	Username     string    `json:"username" form:"username"`
	Password     string    `json:"password" form:"password"`
	Time         time.Time `json:"time" form:"time"`
	Collect      string    `json:"collect"`          //文章收藏
	Communcation string    `json:"communcation"`     //交流，聊天
}
