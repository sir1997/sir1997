package dao

import (
	"fmt"
	"time"
	"todo/model"
)

//查询用户是否被注册

func LogonRepise(name string) *model.LoginStruct{
	 username :="username"
  names := &model.LoginStruct{} 
	err := DB.Where(username +"=?",name).Find(names).Error
	
	if err !=nil{
		return nil
	}
    return names
}


//数据保存到数据库
func LogonUser(name string, password string) {

	logon := model.LoginStruct{
		Username: name,
		Password: password,
		Time:     time.Now(),
	}

	if err := DB.Create(&logon).Error; err != nil {
		fmt.Println(err)
	}

}
