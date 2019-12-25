package main

import (
	_ "github.com/crshi/lmz/routers"
	_ "github.com/crshi/lmz/init"
	
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

