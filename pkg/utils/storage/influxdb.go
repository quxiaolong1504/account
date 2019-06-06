package storage

import (
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/quxiaolong/account/pkg/config"
)

var dbClient *client.Client


func InitInfluxDBClient(conf *config.InfluxDBConfig)  {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: conf.Uri(),
		Username: conf.UserName,
		Password: conf.Password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}

	dbClient = &c
}

func ShutDownInfluxdb() {
	(*dbClient).Close()
}