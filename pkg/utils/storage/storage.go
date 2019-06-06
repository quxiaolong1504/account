package storage

import "github.com/quxiaolong/account/pkg/config"

func InitAllStorage() {
	InitMysql(config.Conf.DataBase)
	InitRedis(config.Conf.Redis)
	InitSessionManager(Cache)
	InitInfluxDBClient(config.Conf.InfluxDB)
}

func ShutDownStorage() {
	ShutDownMysql()
	ShutdownRedis()
	ShutDownInfluxdb()
}