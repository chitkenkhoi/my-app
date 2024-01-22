package routes

import (
	"github.com/gin-gonic/gin"
	"server/repository"
)

func Routes(router *gin.Engine) {
	api := router.Group("/root")
	{
		//api.GET("/countItem/:search", func(ctx *gin.Context) {})
		//api.GET("/getItem/:page/:search", func(ctx *gin.Context) {})
		//api.GET("searchItem/:search", func(ctx *gin.Context) {})
		//api.POST("/createItem", func(ctx *gin.Context) {})
		//api.POST("/updateItem/:id", controller.HandleUpdateItem)
		//api.DELETE("/deleteItem/:id", controller.HandleDeleteItem)
		//api.POST("/createOrder", controller.HandleCreateOrder)
		//api.GET("/getOrder1/:buyer/:page", controller.HandleGetOrder1)
		//api.GET("/getOrder2/:buyer/:page", func(ctx *gin.Context) {})
		//api.GET("/getOrder3/:buyer/:page", func(ctx *gin.Context) {})
		//api.GET("getOrder/:id", controller.HandleGetOrderId)
		//api.GET("/updateOrderStatus/:id/:status", controller.HandleUpdateOrderStatus)
		//api.DELETE("/deleteOrder/:id", controller.HandleDeleteOrder)
		//api.POST("/login", controller.HandleLogin)
		//api.GET("/logout", controller.HandleLogout)
		api.POST("/test", repository.Test)

	}
}
