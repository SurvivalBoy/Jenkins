package service

import (
	"github.com/Jenkins/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReturnSuccessResponse 接口返回数据信息
func ReturnSuccessResponse(ctx *gin.Context, data interface{}) {

	response := &setting.SuccessBody{
		"0",
		data,
	}
	ctx.JSON(http.StatusOK, response)
	return
}

// ReturnErrorResponse 返回错误
func ReturnErrorResponse(ctx *gin.Context, err string, errmsg string) {
	response := &setting.ErrorBody{
		"1",
		errmsg,
		err,
	}
	ctx.JSON(http.StatusOK, response)
	return
}
