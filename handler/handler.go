package handler

import (
	"github.com/gin-gonic/gin"
	. "jenkins_demo/service"
)

// HandleAddUser
func HandleAddUser(ctx *gin.Context) {
	name := ctx.PostForm("name")
	if len(name) == 0 {
		ReturnErrorResponse(ctx, "", "Please provide name.")
		return
	}

	mobile := ctx.PostForm("mobile")
	if len(mobile) == 0 {
		ReturnErrorResponse(ctx, "", "Please provide mobile.")
		return
	}

	if err := AddUser(name, mobile); err != nil {
		ReturnErrorResponse(ctx, err.Error(), "")
		return
	}

	ReturnSuccessResponse(ctx, "add success")
}

// HandleUserInfo
func HandleUserInfo(ctx *gin.Context) {
	name := ctx.PostForm("name")
	if len(name) == 0 {
		ReturnErrorResponse(ctx, "", "Please provide name.")
		return
	}

	user := UserInfo(name)
	ReturnSuccessResponse(ctx, user)
}

// HandleUpdateUser
func HandleUpdateUser(ctx *gin.Context) {
	name := ctx.PostForm("name")
	if len(name) == 0 {
		ReturnErrorResponse(ctx, "", "Please provide name.")
		return
	}
	mobile := ctx.PostForm("mobile")
	if len(mobile) == 0 {
		ReturnErrorResponse(ctx, "", "Please provide mobile.")
		return
	}
	if err := UpdateUser(name, mobile); err != nil {
		ReturnErrorResponse(ctx, err.Error(), "")
		return
	}

	ReturnSuccessResponse(ctx, "update success")
}
