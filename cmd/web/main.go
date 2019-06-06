package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/quxiaolong/account/pkg/config"
	"github.com/quxiaolong/account/pkg/utils"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"github.com/quxiaolong/account/pkg/web/routers"
)


func init(){
	config.Conf.Load("./etc")
	storage.InitAllStorage()
	utils.InitWeChat(config.Conf.WeChat)

}

func main() {
	r := routers.NewRouter()
	r.Run(":8080")
}
