package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/rpcs/seqd"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		id := seqd.Generate()
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("pong id:%d node: %d id:%d", id, id.Node(), id.Time()),
		})
	})
	r.Run()
}
