package router

import (
	"github.com/astaxie/beego"
	"github.com/ericdoug-qi/WeShop/weshop/cloud/controller"
)

func init() {
	beego.Router("/getlist", &controller.ShopController{}, "*:GetProductList")
}
