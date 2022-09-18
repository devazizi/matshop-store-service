package routes

import (
	"net/http"
	"store/adapters"
	v1 "store/delivery/http/v1"

	"github.com/gin-gonic/gin"
)

func RouteV1(router *gin.Engine, db adapters.DB) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotAcceptable, map[string]any{
			"message": "not acceptable request",
		})
	})

	apiV1 := router.Group("/api/v1")

	stores := apiV1.Group("")
	{
		stores.GET("/stores", v1.AdminListStores(db))
		stores.GET("/stores/{id}", v1.AdminShowStore(db))
		stores.GET("/store/me", v1.GetClientStore(db))
		stores.POST("/store", v1.RegisterStore(db))
		stores.PUT("/store/me", v1.UpdateClientStore(db))
		stores.PUT("/store/{id}", v1.UpdateStoreByAdmin(db))
		stores.DELETE("/store/{id}", v1.BanStoreByAdmin(db))
	}

	// productCategories := apiV1.Group("")
	// {
	// 	// 	// productCategories.GET("/product-categories")
	// 	// 	// productCategories.GET("/product-categories/{id}")
	// 	productCategories.GET("/product-categories/public")
	// 	productCategories.GET("/product-categories/{id}/public")
	// 	// 	// productCategories.POST("/product-categories")
	// 	// 	// productCategories.DELETE("/product-categories/{id}")
	// 	// 	// productCategories.PUT("/product-categories/{id}")
	// }

	media := apiV1.Group("")
	{
		media.GET("/media/{modelType}/{modelId}", v1.GetMedia())
		media.POST("/media", v1.StoreMedia())
		media.DELETE("/media/{id}", v1.DestroyMedia())
	}

	// products := v1.Group("")
	// {
	// 	// vendor side
	// 	// products.GET("/products/me")
	// 	// products.GET("/products/{id}/me")
	// 	// products.POST("/products/me")
	// 	// products.DELETE("/products/{id}/me")
	// 	// products.PUT("products/{id}/me")

	// 	// admin side
	// 	// products.GET("/products/admin")
	// 	// products.GET("/products/{id}/admin")
	// 	// products.POST("/products/admin")
	// 	// products.DELETE("/products/{id}/admin")
	// 	// products.PUT("products/{id}/admin")

	// 	// client side
	// 	// products.GET("products")
	// 	// products.GET("/products/search")
	// 	// products.GET("/products/{id}")
	// }

}
