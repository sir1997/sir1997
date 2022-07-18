package dao

import (
	_"todo/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Msql() (err error) {
	sql := "root:1997619@(localhost:3306)/tangyun?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open("mysql", sql)
	if err != nil {
		return err
	}
	//DB.AutoMigrate(&model.LoginStruct{})
	return DB.DB().Ping()
}

