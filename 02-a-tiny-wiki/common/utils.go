package common

import "github.com/gin-gonic/gin"

func RespErr(err string) gin.H {
	return gin.H{
		"message": err,
		"code":    0,
	}
}

func RespOk(obj any) gin.H {
	return gin.H{
		"message": "ok",
		"code":    1,
		"data":    obj,
	}
}
