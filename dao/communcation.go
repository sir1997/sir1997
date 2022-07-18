package dao

import "fmt"

//用户之间实现交流会话，实现一个短暂会话，不保存数据
func Communcation(com string,author string) (string ,string){

		Redis.Set(author,com,0)

		val,err := Redis.Get(author).Result()
		if err !=nil{
			fmt.Println("查询失败")
			
		}
		return val,author

}