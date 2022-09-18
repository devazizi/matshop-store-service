package v1

import (
	"store/adapters"

	"github.com/gin-gonic/gin"
)

func RegisterStore(db adapters.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func AdminListStores(db adapters.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func AdminShowStore(db adapters.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func GetClientStore(db adapters.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateClientStore(db adapters.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UpdateStoreByAdmin(db adapters.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func BanStoreByAdmin(db adapters.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
