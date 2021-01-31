package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/ericdoug-qi/Weshop/weshop/cloud/router"
)

type ShopConf struct {
	redisAddr string
}

var shopConf = &ShopConf{}


func initAppConf() {
	redisAddr := beego.AppConfig.String("redis_addr")
	logs.Debug("Get redis addr from config, the addr is %s", redisAddr)

	shopConf.redisAddr = redisAddr
}

func main() {

	initAppConf()
	beego.Run()
}