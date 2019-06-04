package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/quxiaolong/account/pkg/config"
	"github.com/quxiaolong/account/pkg/models"
	"github.com/quxiaolong/account/pkg/utils/storage"
)


func init(){
	config.Conf.Load("./etc")
	storage.InitDB(config.Conf.DataBase)
}

func main() {
	user  := &models.User{}
	//user := &models.User{
	//	Phone: "+86-18500351504",
	//}
	//storage.Mysql.Master.Create(user)

	storage.Mysql.Slaves[0].Where("id > ?", 1).First(user)
	user.SetPassword("123")
	fmt.Print(user.VerifyPassword("123"))
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": fmt.Sprintf("pong id:%d node: %d id:%d"),
	//	})
	//})
	//r.Run()
}
