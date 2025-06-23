package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}
func (p *ProductHandler) GetProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get List Products",
	})
}
func (p *ProductHandler) GetProductsBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Product By Slug",
		"slug":    slug,
	})
}
func (p *ProductHandler) PostProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post Product",
	})
	
}
func (p *ProductHandler) PutProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update Product",
	})
}
func (p *ProductHandler) DeleteProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Delete Product",
	})
}