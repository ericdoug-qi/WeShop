module github.com/ericdoug-qi/WeShop

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/eddycjy/blog v0.0.0-20201231131640-f550aab68464 // indirect
	github.com/garyburd/redigo v1.6.2 // indirect
	github.com/jmoiron/sqlx v1.3.1 // indirect
)

replace (
	github.com/ericgoug-qi/WeShop/weshop/cloud/router => ./weshop/cloud/router
)
