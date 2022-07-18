package dao

import (
	"todo/model"
	"github.com/astaxie/session"
	"github.com/jinzhu/gorm"
)

var Session session.Session

//插入数据,调试插入文章使用
//前端的结构体字段绑定————传到结构体变量————通过结构体字段传入数据库
func InserArticle(articl *model.ArticleStruct) (db *gorm.DB) {
	sqlr := `insert into 
	articlestruct(
     model.title,model.username,model.content,model.like,model.collect,model.comment) values(?,?,?,?,?,?)`

	db = DB.Exec(sqlr, articl.Title, articl.Username, articl.Content, articl.Like, articl.Collect, articl.Comment)
	if db != nil {
		return
	}

	return nil

}

//传出所有文章数据到前端页面

func Outarticle() (articleLis []*model.ArticleStruct) {

	err := DB.Find(&articleLis).Error
	if err!=nil{
		return nil
	}

	return 

}

//传出指定文章到前端，详情页
func OneArticle(title string) (article []*model.ArticleStruct) {

	err := DB.Where("title=?", title).Find(&article).Error
	if err != nil {
		return nil
	}
	return
}

//分类文章查找
func ClassifyFinde(classi string) (classify []*model.ArticleStruct) {

	err := DB.Where("classify=?", classi).Find(&classify).Error
	if err != nil {
		return nil
	}
	return

}
