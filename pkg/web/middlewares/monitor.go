package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"strings"
)

func Monitor() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// 打 qps 点
		go func() {
			batchPoint,_ := client.NewBatchPoints(client.BatchPointsConfig{
				Database: "monitor",
				Precision: "ns",
			})
			args := strings.Split(c.HandlerName(), ".")
			name := fmt.Sprintf("%s_%s", args[len(args)-1], c.Request.Method)
			fmt.Print()
			point1, _ := client.NewPoint(
				"qps",
				map[string]string{"handler": name, "method": c.Request.Method},
				map[string]interface{} {
					"value": 1,
				},)

			if c.Writer.Status() >= 500 || c.Writer.Status() < 600 {
				p5xx, _ := client.NewPoint(
					"5xx",
					map[string]string{"handler": name, "method": c.Request.Method},
					map[string]interface{} {
						"value": 1,
					},
				)
				batchPoint.AddPoint(p5xx)
			}


			batchPoint.AddPoint(point1)
			_ = storage.InfluxDB.Write(batchPoint)
		}()

		// TODO: 打请求耗时的点 & status code


		c.Next()
	}

	return gin.HandlerFunc(fn)
}
