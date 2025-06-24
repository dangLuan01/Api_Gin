package main

import (
	handler "github.com/dangLuan01/api_gin/internal/api/v1/handler"
	"github.com/dangLuan01/api_gin/utils"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	// Register custom validation
	if err := utils.RegisterValidators(); err != nil {
		panic(err)
	}
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