package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zh-liang-cn/learning-go/02-a-tiny-wiki/common"
	"github.com/zh-liang-cn/learning-go/02-a-tiny-wiki/model"
)

func Index() {

}

func GetPage(c *gin.Context) {
	path := c.Param("path")

	page, err := model.GetPageByPath(path)
	var obj interface{}

	if err != nil {
		obj = common.RespErr(err.Error())
	} else {
		obj = common.RespOk(page)
	}

	c.JSON(http.StatusOK, obj)
}

func SavePage(c *gin.Context) {
	path := c.Param("path")
	title := c.PostForm("title")
	content := c.PostForm("content")

	page := model.SavePage(path, title, content)

	rsp := map[string]any{
		"id": page.ID,
	}

	c.JSON(http.StatusOK, common.RespOk(rsp))
}
