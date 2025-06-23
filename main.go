package main

import (
	"github.com/gin-gonic/gin"
	handler "github.com/dangLuan01/api_gin/internal/api/v1/handler"
)

func main()  {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		userHandler := handler.NewUserHandler()
		user := v1.Group("/users")
		{
			user.GET("", userHandler.GetUsersV1)
			user.GET("/:id", userHandler.GetUserByIdV1)
			user.POST("", userHandler.PostUserV1)
			user.PUT("/:id", userHandler.PutUserV1)
			user.DELETE("/:id", userHandler.DeleteUserV1)

		}
		productHandler := handler.NewProductHandler()
		product := v1.Group("/products")
		{
			product.GET("", productHandler.GetProductsV1)
			product.GET("/:slug", productHandler.GetProductsBySlugV1)
			product.POST("", productHandler.PostProductsV1)
			product.PUT("/:id", productHandler.PutProductsV1)
			product.DELETE("/:id", productHandler.DeleteProductsV1)
		}
	}
	r.Run(":8080")
}