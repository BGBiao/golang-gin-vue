/*================================================================
*Copyright (C) 2019 BGBiao Ltd. All rights reserved.
*
*FileName:main.go
*Author:Xuebiao Xu
*Date:2018年12月02日
*Description:
*
================================================================*/
package main

import (
	"expvar"
	_ "expvar"
	_ "fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to bgbiao's SRE Site, please visit https://bgbiao.top ! Wechat: CloudNativeOps ",
	})
}

func main() {
	r := gin.Default()

	// inner metrics
	r.GET("/debug/vars", gin.WrapH(expvar.Handler()))
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", HelloPage)
		v1.GET("/hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})

		v1.GET("/line", func(c *gin.Context) {
			// 注意:在前后端分离过程中，需要注意跨域问题，因此需要设置请求头
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			legendData := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
			xAxisData := []int{120, 240, rand.Intn(500), rand.Intn(500), 150, 230, 180}
			c.JSON(200, gin.H{
				"legend_data": legendData,
				"xAxis_data":  xAxisData,
			})
		})
	}
	//定义默认路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	r.Run(":8000")
}
