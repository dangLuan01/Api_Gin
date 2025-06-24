package handler

import (
	"net/http"

	"github.com/dangLuan01/api_gin/utils"
	"github.com/gin-gonic/gin"
)
type ProductHandler struct{}

type GetProductBySlugV1Params struct {
	Slug 			string `uri:"slug" binding:"slug,min=1"`
}
type GetProductsV1Params struct {
	Search 			string `form:"search" binding:"required,min=3,max=100"`
	Limit 			int    `form:"limit" binding:"omitempty,min=1,max=100"`
}
type ProductImage struct {
	Img_Name 		string `json:"img_name" binding:"required,min=3,max=100"`
	Image 			string `json:"image" binding:"required,file_ext=jpg jpeg png gif"`
}
type ProductAttributes struct {
	Size 			string `json:"size" binding:"required"`
	Color 			string `json:"color" binding:"required"`
}
type PostProductsV1Params struct {
	Name 				string 							`json:"name" binding:"required,min=3,max=100"`
	Price 				int 							`json:"price" binding:"required,minInt=0,maxInt=2000"`
	Status 				*bool 							`json:"status" binding:"omitempty"`
	ProductImage 		ProductImage 					`json:"product_image" binding:"required"`
	Tags 				[]string 						`json:"tags" binding:"required,gt=0"`
	ProductAttributes 	map[string]ProductAttributes 	`json:"product_attributes" binding:"required,gt=0,dive"`
	ProductMetadata 	map[string]interface{} 			`json:"product_metadata" binding:"required,gt=0,dive"`
}
func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}
func (p *ProductHandler) GetProductsV1(ctx *gin.Context) {
	var params GetProductsV1Params
	
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandlerValidationErrors(err))
		return
	}
	if params.Limit == 0 {
		params.Limit = 1 // Default limit
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get List Products",
		"search":  params.Search,
		"limit":   params.Limit,
	})
}
func (p *ProductHandler) GetProductsBySlugV1(ctx *gin.Context) {
	var params GetProductBySlugV1Params
	if err := ctx.ShouldBindUri(&params); err != nil{
		ctx.JSON(http.StatusBadRequest, utils.HandlerValidationErrors(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Product By Slug",
		"slug":    params.Slug,
	})
}
func (p *ProductHandler) PostProductsV1(ctx *gin.Context) {
	var params PostProductsV1Params
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadGateway, utils.HandlerValidationErrors(err))
		return
	}
	if params.Status == nil {
		defaultStatus := true 
		params.Status = &defaultStatus
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": 				"Post Product",
		"name":    				params.Name,
		"price":   				params.Price,
		"product_image":   		params.ProductImage,
		"status":  				params.Status,
		"tags": 				params.Tags,
		"product_attributes": 	params.ProductAttributes,
		"product_metadata": 	params.ProductMetadata,
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