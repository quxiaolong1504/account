package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/quxiaolong/account/pkg/config"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"github.com/quxiaolong/account/pkg/web/handlers"
	"github.com/quxiaolong/account/pkg/web/routers"
)


func init(){
	config.Conf.Load("./etc")
	storage.InitDB(config.Conf.DataBase)
}

func main() {
	r := routers.NewRouter()
	r.POST("/register", handlers.RegisterHandler{}.Post)
	r.Run(":8080")
}
