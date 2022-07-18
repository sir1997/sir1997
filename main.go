package main

import (
	"log"
	
	"todo/dao"
	"todo/service"
	
)

func main() {

	//链接数据库
	err := dao.Msql()
	if err != nil {
		log.Fatal("suscess")
	}


	r := service.EngineStater()

	r.Run(":8085")
}
