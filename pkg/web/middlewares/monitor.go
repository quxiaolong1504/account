package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"net/http"
	"strings"
	"time"
)

func Monitor() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		start := time.Now()
		// 打 qps 点
		go func() {
			batchPoint,_ := client.NewBatchPoints(client.BatchPointsConfig{
				Database: "monitor",
				Precision: "ns",
			})
			args := strings.Split(c.HandlerName(), ".")
			name := fmt.Sprintf("%s_%s", args[len(args)-1], c.Request.Method)

			// performance point
			point := genPerformancePoint(name, start,  c)

			// status point
			statusPoint := genStatusPoint(name, c)
			if statusPoint != nil {
				batchPoint.AddPoint(statusPoint)
			}

			batchPoint.AddPoint(point)
			err := storage.InfluxDB.Write(batchPoint)
			if err!= nil {
				fmt.Print(err.Error())
			}
		}()

		c.Next()
	}

	return gin.HandlerFunc(fn)
}


func genPerformancePoint(handlerName string, start time.Time, c *gin.Context) *client.Point {
	rt := time.Since(start).Seconds()
	point1, _ := client.NewPoint(
		"performance",
		map[string]string{"handler": handlerName, "method": c.Request.Method},
		map[string]interface{} {
			"qps": 1,
			"rt": rt,
		},)
	return point1
}

func genStatusPoint(name string, c *gin.Context) *client.Point{
	if c.Writer.Status() >= http.StatusUnauthorized && c.Writer.Status() < http.StatusNetworkAuthenticationRequired {
		_5xx := 0
		_401 := 0
		if c.Writer.Status() >= http.StatusInternalServerError && c.Writer.Status() <= http.StatusNetworkAuthenticationRequired{
			_5xx = 1
		}
		if c.Writer.Status() == http.StatusUnauthorized {
			_401 = 1
		}
		point, _ := client.NewPoint(
			"status",
			map[string]string{"handler": name, "method": c.Request.Method},
			map[string]interface{} {
				"5xx": _5xx,
				"401": _401,
			},
		)
		return  point
	}
	return nil
}