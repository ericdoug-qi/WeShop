package controller

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ShopController struct {
	beego.Controller
}

varl DB *sqlx.DB


type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Intro string `json:"intro"`
	Price float64 `json:"price"`
	Price_yh float64 `json:"price_yh"`
	Photo_little string `json:"photo_little"`
	Saledcount int `json:"saledcount"`
}

func init() {
	database, err := sqlx.Open("mysql",
			"miniprogram:miniprogram@tecp(127.0.0.1:3306)")

	if err != nil {
		fmt.Println("Open MySQL error: ", err)
	}

	DB = database

}

func (this *ShopController) GetpRoductList() {
	cat_id, err := this.GetInt("cat_id")
	if err != nil {
		fmt.Println("Get cat_id failed, err: ", err)
		return
	}

	var product []Product
	err = DB.Select(&product, "select id, name, intro, price, price_yh, price_little, saledcount from `miniprogram_product` where cid=?", cat_id)
	if err != nil {
		fmt.Println("Select data from product failed, err: ", err)
		return
	}

	this.Data["json"] = product

	this.ServeJSON()


}