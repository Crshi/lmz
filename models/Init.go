package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init初始化
func init() {
	orm.RegisterModel(new(Article))
}

// 表名
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

// Article表名
func ArticleTBName() string {
	return TableName("Article")
}
