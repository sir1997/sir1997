package dao

import "fmt"

//收藏别人的文章或其他
func Collect(title string, user string){

	//通过截取标题的一部分作为字段
	key := []byte(title)
	keyrep := key[0:3]
	maps := make(map[string]interface{})
	maps[string(keyrep)] = title

	_,err := Redis.HMSet(user, maps).Result()
	if err !=nil{
		fmt.Println("创建失败")
	}

	
}

func CollectShow(user string)[]string{
	val, err := Redis.HVals(user).Result()
	if err != nil {
		fmt.Println("查询失败")
	}
	return val
}