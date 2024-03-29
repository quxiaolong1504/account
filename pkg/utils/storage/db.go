package storage


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/quxiaolong/account/pkg/config"
	"github.com/quxiaolong/account/pkg/utils/logger"
)


type MysqlCli struct {
	Master *gorm.DB
	Slaves []*gorm.DB
}

var Mysql *MysqlCli

func InitDB(conf *config.DBConfig) {
	master , err := gorm.Open(conf.Master.Engine,  conf.Master.Uri())
	master.DB().SetMaxOpenConns(int(conf.Master.MaxOpenConns))
	if err != nil {
		logger.Logger.Critical(err.Error())
	}

	slaves := make([]*gorm.DB, 0)
	for _, slave := range(conf.Slaves){
		cli , err := gorm.Open(slave.Engine, slave.Uri())
		if err != nil {
			logger.Logger.Critical(err.Error())
		}
		cli.DB().SetMaxOpenConns(int(slave.MaxOpenConns))
		slaves =  append(slaves, cli)
	}

	Mysql = &MysqlCli{
		Master: master,
		Slaves: slaves,
	}
}

func ShutDown() {
	Mysql.Master.Close()
	for i:=0; i <= len(Mysql.Slaves); i++ {
		Mysql.Slaves[i].Close()
	}
}