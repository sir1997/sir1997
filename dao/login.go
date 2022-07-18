package dao

import (
	"todo/model"

)

//查询数据登录认证
func SetLogin(name string) (log []*model.LoginStruct) {
	//DB.Select("where=id?",name,password)
    
	username := "username"
	_ = DB.Where(username+"=?",name).Find(&log)
	
	return 
}
