package models

import (
	"github.com/astaxie/beego/orm"
)

//Init初始化
func Init() {
	orm.RegisterModel(new(Article))
}
