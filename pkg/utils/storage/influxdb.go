package storage

import (
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/quxiaolong/account/pkg/config"
	"github.com/quxiaolong/account/pkg/utils/logger"
)

var InfluxDB client.Client


func InitInfluxDBClient(conf *config.InfluxDBConfig)  {
	logger.Logger.Infof("Init InfluxDB client")
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: conf.Uri(),
		Username: conf.UserName,
		Password: conf.Password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	InfluxDB = c

}

func ShutDownInfluxDB() {
	InfluxDB.Close()
}