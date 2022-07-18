package dao

import (
	"fmt"
	"strconv"
	"todo/model"

	"github.com/go-redis/redis"
)

var i int = 0  //全局浏览自增，避免操作数据库造成压力
//浏览的实现，数据库字段的更新，也可以使用redis，以下方法搭建接口即可
func ViewCount(title string) error{

	if title !=""{
		i++
	article := model.ArticleStruct{}
   err := DB.Model(&article).Where("title=?",title).Update("count",i).Error

      if err!=nil{
		return err
	  }
	}
	return nil
}

//点赞功能的实现，操作数据库字段
func LikeCount(title string,like string)error{

	likeInt,err:= strconv.Atoi(like)
	if err!=nil{
		fmt.Println("转换失败")
	}
   
		article := model.ArticleStruct{}
		err = DB.Model(&article).Where("title=?",title).Update("like",likeInt+1).Error
		
		if err!=nil{
			return err
		  }
	
	return nil
}













//执行点赞与取消点赞
func ActionLike(client *redis.Client, userid string) {
	val, err := client.SIsMember("like_id", userid).Result()
	if err != nil {
		panic(err)
	}

	//添加点赞，redis表里面创建值
	if val {
		_, err := client.SAdd("like_id", userid).Result()
		if err != nil {
			panic(err)
		}
	} else {
		//删除点赞
		_, err := client.SRem("like_id", userid).Result()
		if err != nil {
			panic(err)
		}
	}
}

//获取点赞次数
func LikeCount_Redis(client *redis.Client) int64 {
	val, err := client.SCard("like_id").Result()
	if err != nil {
		panic(val)
	}
	
	return val
}


func ActionViewCount(client *redis.Client, userid string) {
	
	//添加浏览次数，redis表里面创建值
	if true {
		_, err := client.SAdd("view_count", userid).Result()
		if err != nil {
			panic(err)
		}
	} 
}

//获取浏览数
func ViewCountTatial(client *redis.Client) int64 {
	val, err := client.SCard("view_count").Result()
	if err != nil {
		panic(val)
	}
	
	return val
}