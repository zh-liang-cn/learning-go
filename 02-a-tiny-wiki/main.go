package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zh-liang-cn/learning-go/02-a-tiny-wiki/common"
	"github.com/zh-liang-cn/learning-go/02-a-tiny-wiki/handler"
	"github.com/zh-liang-cn/learning-go/02-a-tiny-wiki/model"
)

func initDBAndMigrate() {
	db := common.InitDB()
	db.AutoMigrate(&model.Page{})
}

func main() {
	initDBAndMigrate()

	r := gin.Default()

	w := r.Group("/w")
	{
		w.POST("/:path", handler.SavePage)
		w.GET("/:path", handler.GetPage)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
