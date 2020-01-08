package init

import (
	"github.com/astaxie/beego"
	"github.com/crshi/lmz/utils"
)

func init() {
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//初始化日志
	utils.InitLogs()
	//初始化缓存
	utils.InitCache()
	//初始化数据库
	InitDatabase()
}
