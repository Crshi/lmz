package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init初始化
func init() {
	orm.RegisterModel(new(Article), new(Tag), new(Catalog), new(TagArticle), new(CatalogArticle))
}

// TableName 表名
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

// ArticleTBName 表名
func ArticleTBName() string {
	return TableName("Article")
}

// TagTBName 表名
func TagTBName() string {
	return TableName("Tag")
}

// CatalogTBName 表名
func CatalogTBName() string {
	return TableName("Catalog")
}

// TagArticleTBName 表名
func TagArticleTBName() string {
	return TableName("TagArticle")
}

// CatalogArticleTBName 表名
func CatalogArticleTBName() string {
	return TableName("CatalogArticle")
}
