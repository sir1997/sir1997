package model

import "time"

type ArticleStruct struct {
	Id       int64     `json:"id"`
	Status   bool      `json:"status"`
	Title    string    `json:"title"`
	Username string    `json:"username"`
	Content  string    `json:"content"`   //文章内容
	Time     time.Time `json:"time"`      //创建时间
	Like     int       `json:"like"`      //点赞
	Collect  int       `json:"collect"`   //收藏
	Comment  string    `json:"comment"`   //评论
	Count    uint32    `json:"count"`     //浏览次数
	Classify string    `json:"classify"`  //分类文章
}


