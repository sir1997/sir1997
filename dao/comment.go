package dao

import (
	"time"
	"todo/model"
)

//加载评论，实现评论功能
func Comment(username string,comment string) (comt []*model.Comment){
	if username!=""&&comment!=""{
		comm :=&model.Comment{Time: time.Now(),Username: username,Comment: comment}

	err := DB.Create(comm).Error
	if err !=nil{
		return nil
	}
	}
	
	commodel :=&model.Comment{}

	err := DB.Model(commodel).Find(&comt).Error
	if err !=nil{
		return nil
	}
	return
}

func CommnetCount() int{
	commodel :=&model.Comment{}
	count := 0
	DB.Model(commodel).Count(&count)
	return count
}