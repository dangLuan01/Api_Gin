package handler

import (
	"net/http"

	"github.com/dangLuan01/api_gin/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}
type GetUserByIdV1Params struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get List Users",
	})
}
func (u *UserHandler) GetUserByUUIdV1(ctx *gin.Context) {
	var params GetUserByIdV1Params
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandlerValidationErrors(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get User By UUID",
		"uuid":      params.Uuid,
	})
}
func (u *UserHandler) PostUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post User",
	})
}
func (u *UserHandler) PutUserV1(ctx *gin.Context) {
	var params GetUserByIdV1Params
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandlerValidationErrors(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Updated User",
		"id":      params.Uuid,
	})
}
func (u *UserHandler) DeleteUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Deleted User",
	})
}