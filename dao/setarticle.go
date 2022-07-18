package dao

import (
	"fmt"
	"time"
	"todo/model"
	"github.com/jinzhu/gorm"
)


//dao.SetArticle(title,str,content)
func SetArticle(title string, author string ,content string,classify string) *gorm.DB {

	 article := &model.ArticleStruct{Title: title, Username: author,Content: content, Time: time.Now(),Classify: classify}
	
	db := DB.Create(&article)
	if db == nil {
		fmt.Println("NILL")
		return nil
	}
	return db
}
